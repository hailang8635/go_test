//package chapter_0209
package main

import (
    "fmt"
    "strings"
    "time"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    fmt.Println(time.Now())
    fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
    fmt.Println(longestPalindrome("lphbehiapswjudnbcossedgioawodnwdruaaxhbkwrxyzaxygmnjgwysafuqbmtzwdkihbwkrjefrsgjbrycembzzlwhxneiijgzidhngbmxwkhphoctpilgooitqbpjxhwrekiqupmlcvawaiposqttsdgzcsjqrmlgyvkkipfigttahljdhtksrozehkzgshekeaxezrswvtinyouomqybqsrtegwwqhqivgnyehpzrhgzckpnnpvajqevbzeksvbezoqygjtdouecnhpjibmsgmcqcgxwzlzztdneahixxhwwuehfsiqghgeunpxgvavqbqrelnvhnnyqnjqfysfltclzeoaletjfzskzvcdwhlbmwbdkxnyqappjzwlowslwcbbmcxoiqkjaqqwvkybimebapkorhfdzntodhpbhgmsspgkbetmgkqlolsltpulgsmyapgjeswazvhxedqsypejwuzlvegtusjcsoenrcmypexkjxyduohlvkhwbrtzjnarusbouwamazzejhnetfqbidalfomecehfmzqkhndpkxinzkpxvhwargbwvaeqbzdhxzmmeeozxxtzpylohvdwoqocvutcelgdsnmubyeeeufdaoznxpvdiwnkjliqtgcmvhilndcdelpnilszzerdcuokyhcxjuedjielvngarsgxkemvhlzuprywlypxeezaxoqfges"))
    fmt.Println(longestPalindrome("flsuqzhtcahnyickkgtfnlyzwjuiwqiexthpzvcweqzeqpmqwkydhsfipcdrsjkefehhesubkirhalgnevjugfohwnlhbjfewiunlgmomxkafuuokesvfmcnvseixkkzekuinmcbmttzgsqeqbrtlwyqgiquyylaswlgfflrezaxtjobltcnpjsaslyviviosxorjsfncqirsjpkgajkfpoxxmvsyynbbovieoothpjgncfwcvpkvjcmrcuoronrfjcppbisqbzkgpnycqljpjlgeciaqrnqyxzedzkqpqsszovkgtcgxqgkflpmrikksaupukdvkzbltvefitdegnlmzeirotrfeaueqpzppnsjpspgomyezrlxsqlfcjrkglyvzvqakhtvfmeootbtbwfhqucbnuwznigoyatvkocqmbtqghybwrhmyvvuchjpvjckiryvjfxabezchynfxnpqaeampvaapgmvoylyutymdhvhqfmrlmzkhuhupizqiujpwzarnszrexpvgdmtoxvjygjpmiadzdcxtggwamkbwrkeplesupagievwsaaletcuxtpsxmbmeztcylsjxvhzrqizdmgjfyftpzpgxateopwvynljzffszkzzqgofdlwyknqfruhdkvmvrrjpijcjomnrjjubfccaypkpfokohvkqndptciqqiscvmpozlyyrwobeuazsawtimnawquogrohcrnmexiwvjxgwhmtpykqlcfacuadyhaotmmxevqwarppknoxthsmrrknu"))
    fmt.Println(longestPalindrome("aaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaa"))
    fmt.Println(longestPalindrome("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffgggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"))
    fmt.Println(longestPalindrome("1234a4321"))
    fmt.Println(longestPalindrome("ababababa"))
    fmt.Println(longestPalindrome("xaabacxcabaaxcabaax"))
    fmt.Println(longestPalindrome("aacabdkacaa"))
    fmt.Println(longestPalindrome("cbbd"))
    fmt.Println(longestPalindrome("a"))
    fmt.Println(longestPalindrome("ccc"))
    fmt.Println(longestPalindrome("babad"))
    fmt.Println(longestPalindrome("aaabaaaa"))
    fmt.Println(longestPalindrome("ababababababa"))
    fmt.Println(longestPalindrome("abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa"))
    fmt.Println(time.Now())
}

func longestPalindrome(s string) string {
    reverseStr := reverseFunc(s)

    length := len(s)

    maxLength := 0
    var result string

    bound := length/2
    if length%2 != 0 {
        bound += 1
    }

    tmpStrMap := make(map[string]int)

    count := 0

    for i:=0; i<=length-1; i++ {
        //for j:=i+1; (j<=length && j-i <= bound+2); j++ {
        for j:=i+1; (j<=length); j++ {
            if j-i + 3 < maxLength {
                //fmt.Println(maxLength, j, i, j-i)
                continue
            }

            count++
            str := s[i:j]

            //_, val := tmpStrMap[str]
            //if val {
            //    //continue
            //}

            // 正反均存在的字符串
            if strings.Contains(reverseStr, str) {
                strRe := reverseFunc(str)

                //fmt.Println(i, j, str)
                if str == strRe {
                    //fmt.Println(" --> hit: ", str, strRe, reverseStr)
                    if len(str) > maxLength {
                        maxLength = j - i
                        result = str
                    }
                }
                if j+1 <= len(s) {

                    var strTmp string
                    x := s[j: j+1]

                    if strings.Contains(s, str+strRe) {
                        strTmp = str + strRe
                    }
                    if strings.Contains(s, str + x + strRe) {
                        strTmp = str + x + strRe
                    }

                    if len(strTmp) > maxLength {
                       maxLength = len(strTmp)
                       result = strTmp
                    }

                }

                tmpStrMap[result] = maxLength
            } else {
                break
                //fmt.Println(str)
            }
        }

        if maxLength >= length-i {
            break
        }
    }
    fmt.Println(" - ", count)
    return result
}

func longestPalindrome_timeout(s string) string {
    reverseStr := reverseFunc(s)

    length := len(s)

    maxLength := 0
    var result string

    for i:=0; i<=length-1; i++ {
        for j:=i+1; j<=length; j++ {
            str := s[i:j]
            if strings.Contains(reverseStr, str) {
                strRe := reverseFunc(str)
                if str == strRe || strings.Contains(s, str+strRe) {
                    //fmt.Println(" --> hit: ", str, strRe, reverseStr)
                    //fmt.Println(strings.Contains(s, str + ))
                    if len(str) > maxLength {
                        maxLength = j - i
                        result = str
                    }
                }
            }
        }
    }
    return result
}

func longestPalindromeBest(s string) string {
    //双指针法
    start := 0//记录最长回文子串的起始位置
    end := 0 //记录最长回文子串的末尾位置
    ln := len(s)//s的长度
    for i := 0; i < ln; {
        //左右指针，向两边扩展
        l, r := i, i

        //先考虑回文子串可能是偶数的情况，让r先走
        for r < ln -1 && s[r] == s[r+1] { //防止越界，一直向后比较，不相等停止
            r++
        }
        //i到达r所扩展的最远长度的下一个字符
        i = r+1

        //两边一起扩展
        for l > 0 && r < ln-1 && s[l-1] == s[r+1] {
            l--
            r++
        }

        //更新最长回文子串的长度和起始点
        if end -start < r - l {
            start = l
            end = r
        }

    }

    return s[start : end+1]
}

func reverseFunc(s string) string{
    length := len(s)
    var reverseStr string
    for i:=length-1; i>=0; i-- {
        reverseStr = reverseStr + s[i:i+1]
    }
    return reverseStr
}


func testMergeLists() {
    // l1 = [1,2,4], l2 = [1,3,4]
    listNode1d := ListNode{6, nil}
    listNode1c := ListNode{5, &listNode1d}
    listNode1b := ListNode{2, &listNode1c}
    listNode1a := ListNode{1, &listNode1b}

    listNode2d := ListNode{4, nil}
    listNode2c := ListNode{4, &listNode2d}
    listNode2b := ListNode{2, &listNode2c}
    listNode2a := ListNode{1, &listNode2b}

    fmt.Println("1 a b c --> ", listNode1a.Val, listNode1b.Val, listNode1c.Val)
    fmt.Println("2 a b c --> ", listNode2a.Val, listNode2b.Val, listNode2c.Val)

    //fmt.Println(MergeTwoLists(&listNode1a, &listNode2a))


    node1 := ListNode{1, nil}
    node2 := ListNode{2, nil}
    fmt.Println(MergeTwoLists(&node1, &node2))



    //fmt.Println("a next next --> ", listNode1a.Val, listNode1a.Next.Val, listNode1a.Next.Next.Val)
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }

    var head ListNode
    var tail *ListNode

    tail = &head

    currentNode1 := list1
    currentNode2 := list2

    for true {
        var currentNode ListNode

        if currentNode1 == nil && currentNode2 == nil {
            break
        } else if (currentNode1 != nil && currentNode2 == nil) || (currentNode1 != nil && currentNode2 != nil && currentNode1.Val < currentNode2.Val) {
            currentNode.Val = currentNode1.Val
            if currentNode1.Next != nil {
                currentNode1 = currentNode1.Next
                //fmt.Println(currentNode1.Val, currentNode1.Next)
            } else {
                currentNode1 = nil
            }
        } else if (currentNode1 == nil && currentNode2 != nil) || (currentNode1 != nil && currentNode2 != nil && currentNode1.Val >= currentNode2.Val) {
            currentNode.Val = currentNode2.Val
            if currentNode2.Next != nil {
                currentNode2 = currentNode2.Next
                //fmt.Println(currentNode2.Val, currentNode2.Next)
            } else {
                currentNode2 = nil
            }
        } else {
            fmt.Println("exception")
        }

        tail.Next = &currentNode
        tail = &currentNode

        //fmt.Println(tail, currentNode, currentNode1, currentNode2)
    }

    head = *(head.Next)
    ShowNode(&head)
    return &head
}



func ShowNode(p *ListNode) {
    fmt.Println("-------show start -- ")
    for p != nil {
        fmt.Println(" --> ", *p)
        p = p.Next
    }
    fmt.Println("-------show end --")
}