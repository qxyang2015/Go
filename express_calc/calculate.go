package express_calc

import (
	"fmt"
	"unicode"
)

/*
从左到右逐个字符遍历中缀表达式。
输出的字符序列即是后缀表达式：遇到数字直接输出。
遇到运算符则判断：
	1.栈顶运算符优先级更低则入栈，更高或相等则直接输出
	2.栈为空、栈顶是 ( 直接入栈
	3.运算符是 ) 则将栈顶运算符全部弹出，直到遇见 )
中缀表达式遍历完毕，运算符栈不为空则全部弹出，依次追加到输出。
*/

func infix2ToPostfix(exp string) string {
	stack := stack.ItemStack{}
	postfix := ""
	expLen := len(exp)

	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])

		switch char {
		case " ":
			continue
		case "(":
			// 左括号直接入栈
			stack.Push("(")
		case ")":
			// 右括号则弹出元素直到遇到左括号
			for !stack.IsEmpty() {
				preChar := stack.Top()
				if preChar == "(" {
					stack.Pop() // 弹出 "("
					break
				}
				postfix += preChar
				stack.Pop()
			}

			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1

		default:
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for !stack.IsEmpty() {
				top := stack.Top()
				if top == "(" || isLower(top, char) {
					break
				}
				postfix += top
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}

	// 栈不空则全部输出
	for !stack.IsEmpty() {
		postfix += stack.Pop()
	}

	return postfix
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

func pre2post(exps []JsonExpr) ([]JsonExpr, error) {
	log4sys := util.GetLogger()
	var exps2 []JsonExpr
	stack := util.NewStack()
	for _, exp := range exps {
		s := exp.Value
		if exp.Type == FIELD_TYPE_OPERATOR {
			if s == "(" {
				stack.Push(exp)
			} else if s == ")" {
				for {
					expTemp := stack.Pop().(JsonExpr)
					if expTemp.Value == "(" {
						break
					}
					exps2 = append(exps2, expTemp)
				}
			} else {
				//操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
				for !stack.Empty() {
					topVal := stack.Peak().(JsonExpr).Value
					if topVal == "(" {
						break
					}
					if _, ok1 := opInfo[exp.Value]; !ok1 {
						log4sys.Error("opInfo exp.Value[%s] is not found", exp.Value)
						return nil, errors.New(fmt.Sprintf("opInfo exp.Value[%s] is not found", exp.Value))
					}
					if _, ok2 := opInfo[topVal]; !ok2 {
						log4sys.Error("opInfo topVal[%s] is not found", topVal)
						return nil, errors.New(fmt.Sprintf("opInfo topVal[%s] is not found", topVal))
					}
					if opInfo[exp.Value].Salience > opInfo[topVal].Salience {
						break
					}
					exps2 = append(exps2, stack.Pop().(JsonExpr))
				}
				stack.Push(exp)
			}
		} else if exp.Type == FIELD_TYPE_CONSTANT || exp.Type == FIELD_TYPE_VARIABLE || exp.Type == FIELD_TYPE_FUNC {
			exps2 = append(exps2, exp)
		} else {
			log4sys.Error("fieldType[%s] not operator,constant,variable,func", exp.Type)
			return nil, errors.New(fmt.Sprintf("fieldType[%s] not operator,constant,variable,func", exp.Type))
		}
	}
	for {
		if stack.Empty() {
			break
		}
		exps2 = append(exps2, stack.Pop().(JsonExpr))
	}
	return exps2, nil
}
