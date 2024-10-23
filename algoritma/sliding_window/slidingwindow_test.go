package slidingwindow

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	s := "abcabcbb"
	// s := "bbbbb"
	// s := "pwwkew"
	res := slideString(s)
	fmt.Printf("res: %v\n", res)
}

/*
Longest Substring Without Repeating Character ( Leetcode #3)
Given a string s, find the length of the longest substring without repeating characters.

### Dynamic Sliding Windows ####
*/

func slideString(s string) int {
	count := 0
	left := 0
	temp := make(map[string]int, 0)

	for right := range len(s) {
		key := string(s[right])
		if _, ok := temp[key]; ok {
			delete(temp, key)
			left++
		}

		window := (right - left) + 1
		count = max(count, window)
		temp[key]++
	}

	return count
}

/*
Maximum Average Sub Array 1 (LeetCode #643)
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

### Fixed Sliding Windows ###
*/
func TestSlide(t *testing.T) {
	// k := 3
	// nums := []int{2, 1, 5, 1, 3, 2}
	// k := 93
	// nums := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	// k := 7
	// nums := []int{493, 593, 1446, -6013, 2163, 8450, 3008, -1328, 6195, 4102, -6242, -9971, -8327, 4480, -4972, -8305, -1644, 2320, 331, 2465, 2955, -8386, -7580, 1759, -9576, -8088, -9446, -2916, -3600, 923, 1412, -5390, 4492, 9458, -9336, -4754, -3455, 9597, -324, -5628, 4242, 4076, 8159, -2423, -3449, 6744, 9029, -9231, 2283, 6118, -200, 3610, -7566, -6976, 2519, 9532, 2221, -5167, -2370, 1861, -1558, -9815, -1186, 2021, 7050, -4504, 4926, 8362, 7805, -8274, -351, 5826, 7623, -5520, -2395, -8466, -8461, 3261, -3197}

	k := 4
	nums := []int{1, 12, -5, -6, 50, 3}
	resFirst := slideWindow(nums, k)
	resSecond := slideWindowsAgain(nums, k)
	fmt.Printf("res First: %v\n", resFirst)
	fmt.Printf("res Second: %v\n", resSecond)

}

func slideWindowsAgain(nums []int, k int) float64 {
	count := 0
	maxAverage := -9999999
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if i >= k-1 {
			maxAverage = max(maxAverage, count)
			count -= nums[i-(k-1)]
		}
	}
	return float64(maxAverage) / float64(k)
}

func slideWindow(nums []int, k int) float64 {
	// contoh nums = [2, 1, 5, 1, 3, 2], k = 3
	left := 0
	current_sum := 0
	maxAverage := -9999999

	// membuat window sepanjang k
	// isi window berisi w := 2 + 1 + 5
	for left < k {
		current_sum += nums[left]
		left++
	}

	// geser window per 1 elemen dan lakukan kalkulasi.
	// geser elemen ke kanan, left akan ditambah dan right bertambah juga
	// idx   : 0  1  2  3  4  5
	// nums  : 2, 1, 5, 1, 3, 2
	// window: 2, 1, 5     == > geser ke kanan
	//     >>:    1, 5, 1  == > saat geser, hasil kalkukasi kurangi dengan angka 2 dan tambah dengan angka 1

	// karena window sudah terbentuk sampai dengan left atau k,
	// maka untuk geser dimulai dari posisi terakhir window
	for right := left; right < len(nums)-1; right++ {
		// geser ke kanan
		current_sum += nums[right]
		// karena sudah geser ke kanan, maka sebelah kiri dikurangi
		// untuk maintain panjang window ( k )
		current_sum -= nums[right-k]

		maxAverage = max(maxAverage, current_sum)
	}
	return float64(maxAverage) / float64(k)
}

/*
Given two strings s and t of lengths m and n respectively, return the minimum window
substring of s such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
*/

func TestMinWindowSubstring(test *testing.T) {
	// s := "ADOBECODEBANC"
	// t := "ABC"
	// jawaban := "BANC"

	s := "obzcopzocynyrsgsarijyxnkpnukkrvzuwdjldxndmnvevpgmxrmvfwkutwekrffnloyqnntbdohyfqndhzyoykiripdzwiojyoznbtogjyfpouuxvumtewmmnqnkadvzrvouqfbbdiqremqzgevkbhyoznacqwbhtrcjwfkzpdstpjswnpiqxjhywjanhdwavajrhwtwzlrqwmombxcaijzevbtcfsdcuovckoalcseaesmhrrizcjgxkbartdtotpsefsrjmvksqyahpijsrppdqpvmuocofuunonybjivbjviyftsyiicbzxnwnrmvlgkzticetyfcvqcbjvbufdxgcmesdqnowzpshuwcseenwjqhgsdlxatamysrohfnixfprdsljyyfhrnnjsagtuihuczilgvtfcjwgdhpbixlzmakebszxbhrdibpoxiwztshwczamwnninzmqrmpsviydkptjzpktksrortapgpxwojofxeasoyvyprjoguhqobehugwdvtzlenrcttuitsiijswpogicjolfxhiscjggzzissfcnxnvgftxvbfzkukqrtalvktdjsodmtgzqtuyaqvvrbuexgwqzwduixzrpnvegddyyywaquxjxrnuzlmyipuqotkghfkpknqinoidifnfyczzonxydtqroazxhjnrxfbmtlqcsfhshjrxwqvblovaouxwempdrrplefnxmwrwfjtebrfnfanvvmtbzjesctdgbsfnpxlwihalyiafincfcwgdfkvhebphtxukwgjgplrntsuchyjjuqozakiglangxkttsczhnswjksnuqwflmumpexxrznzwxurrysaokwxxqkrggytvsgkyfjrewrcvntomnoazmzycjrjrqemimyhriyxgrzcfuqtjhvjtuhwfzhwpljzajitrhryaqchnuawbxhxrpvyqcvhpggrpplhychyulijhkglinibedauhvdydkqszdbzfkzbvhldstocgydnbfjkcnkfxcyyfbzmmyojgzmasccaahpdnzproaxnexnkamwmkmwslksfpwirexxtymkmojztgmfhydvlqtddewjvsrmyqjrpycbmndhupmdqqabiuelacuvxnhxgtpvrtwfgzpcrbhhtikbcqpctlxszgpfbgcsbaaiapmtsucocmpecgixshrrnhyrpalralbccnxvjzjllarqhznzghswqsnfuyywmzbopyjyauknxddgdthlabjqtwxpxwljvoxkpjjpfvccyikbbrpdsyvlxscuoofkecwtnfkvcnzbxkeabtdusyhrqklhaqreupakxkfzxgawqfwsaboszvlshwzhosojjotgyagygguzntrouhiweuomqptfjjqsxlbylhwtpssdlltgubczxslqjgxuqnmpynnlwjgmebrpokxjnbiltvbebyytnnjlcwyzignmhedwqbfdepqakrelrdfesqrumptwwgifmmbepiktxavhuavlfaqxqhreznbvvlakzeoomykkzftthoemqwliednfsqcnbexbimrvkdhllcesrlhhjsspvfupxwdybablotibypmjutclgjurbmhztboqatrdwsomnxnmocvixxvfiqwmednahdqhxjkvcyhpxxdmzzuyyqdjibvmfkmonfxmohhshpkhmntnoplphqyprveyfsmsxjfosmicdsjrieeytpnbhlsziwxnpmgoxneqbnufhfwrjbqcsdfarybzwaplmxckkgclvwqdbpumsmqkswmjwnkuqbicykoisqwoootrdpdvcuiuswfqmrkctsgrevcxnyncmivsxbpbxzxpwchiwtkroqisnmrbmefbmatmdknaklpgpyqlsccgunaibsloyqpnsibwuowebomrmcegejozypjzjunjmeygozcjqbnrpakdermjcckartbcppmbtkhkmmtcngteigjnxxyzaibtdcwutkvpwezisskfaeljmxyjwykwglqlnofhycwuivdbnpintuyhtyqpwaoelgpbuwiuyeqhbvkqlsfgmeoheexbhnhutxvnvfjwlzfmvpcghiowocdsjcvqrdmkcizxnivbianfpsnzabxqecinhgfyjrjlbikrrgsbgfgyxtzzwwpayapfgueroncpxogouyrdgzdfucfrywtywjeefkvtzxlwmrniselyeodysirqflpduvibfdvedgcrzpzrunpadvawfsmmddqzaaahfxlifobffbyzqqbtlcpquedzjvykvarayfldvmkapjcfzfbmhscdwhciecsbdledspgpdtsteuafzbrjuvmsfrajtulwirzagiqjdiehefmfifocadxfuxrpsemavncdxuoaetjkavqicgndjkkfhbvbhjdcygfwcwyhpirrfjziqonbyxhibelinpllxsjzoiifscwzlyjdmwhnuovvugfhvquuleuzmehggdfubpzolgbhwyeqekzccuypaspozwuhbzbdqdtejuniuuyagackubauvriwneeqfhtwkocuipcelcfrcjcymcuktegiikyosumeioatfcxrheklookaqekljtvtdwhxsteajevpjviqzudnjnqbucnfvkybggaybebljwcstmktgnipdyrxbgewqczzkaxmeazpzbjsntltjwlmuclxirwytvxgvxscztryubtjweehapvxrguzzsatozzjytnamfyiitreyxmanhzeqwgpoikcjlokebksgkaqetverjegqgkicsyqcktmwjwakivtsxjwrgakphqincqrxqhzbcnxljzwturmsaklhnvyungjrxaonjqomdnxpnvihmwzphkyuhwqwdboabepmwgyatyrgtboiypxfavbjtrgwswyvcqhzwibpisydtmltbkydhznbsvxktyfxopwkxzbftzknnwipghuoijrbgqnzovxckvojvsqqraffwowfvqvfcmiicwitrhxdeombgesxexedlakitfovtydxunqnwqqdeeekiwjnwoshqcsljiicgobbbuqakjdonjawgjlezdnqhfdqnmsuavxdpnfzwipmspiabveaarshzwxmirgkmfncvtdrdvfxkpxlkdokxgtwcskmjryyymcthfnkasinihaunohkxaibtsqelockaefjmsuolebtnepauwmrxutspjwaxbmahsjtkfkxlnszribmeofbkyvbjscjtqjakuwvcgunvnonvqbbggfshauqsyznokqbhowjusypfnecffenojfvlblgzntqzlrgzprvhqnpfrrkzxznieiuivajivzijsqijigtatifmbplzqahuidegfoobpymkputzamzvweiyvvzlwihgmmmrcburbgbsdxrfjsbiylitghgcpqjbevvgypxcybubyoijijrhuzcdijfybqbfowlookqmlnplbxvjjosfqviygqyhgamuwzjklbyzopkrnhbywtfoqomweldmlrhjqswctubiknzzvcztyehouvnyiqnvkufaobehxhrjvtisxjlxoumipzjarwvbsaegdkpbsjmpevjbewzuqnfhoohhmdjgfpmjzdmtmykqvtucptwfidpwtwffzolffzqfdearclkyeecuzabjeqhxpmfodsvisnpxrqowdawheydfyhoexvcmihdlzavtqlshdhdgjzpozvvackebhgqppvcrvymljfvooauxcjnbejdivikcoaugxwzsulgfqdtefpehbrlhaoqxwcancuvbqutnfbuygoemditeagmcveatgaikwflozgdhkyfqmjcruyyuemwbqwxyyfiwnvlmbovlmccaoguieu"
	t := "cjgamyzjwxrgwedhsexosmswogckohesskteksqgrjonnrwhywxqkqmywqjlxnfrayykqotkzhxmbwvzstrcjfchvluvbaobymlrcgbbqaprwlsqglsrqvynitklvzmvlamqipryqjpmwhdcsxtkutyfoiqljfhxftnnjgmbpdplnuphuksoestuckgopnlwiyltezuwdmhsgzzajtrpnkkswsglhrjprxlvwftbtdtacvclotdcepuahcootzfkwqhtydwrgqrilwvbpadvpzwybmowluikmsfkvbebrxletigjjlealczoqnnejvowptikumnokysfjyoskvsxztnqhcwsamopfzablnrxokdxktrwqjvqfjimneenqvdxufahsshiemfofwlyiionrybfchuucxtyctixlpfrbngiltgtbwivujcyrwutwnuajcxwtfowuuefpnzqljnitpgkobfkqzkzdkwwpksjgzqvoplbzzjuqqgetlojnblslhpatjlzkbuathcuilqzdwfyhwkwxvpicgkxrxweaqevziriwhjzdqanmkljfatjifgaccefukavvsfrbqshhswtchfjkausgaukeapanswimbznstubmswqadckewemzbwdbogogcysfxhzreafwxxwczigwpuvqtathgkpkijqiqrzwugtr"
	jawaban := "twxpxwljvoxkpjjpfvccyikbbrpdsyvlxscuoofkecwtnfkvcnzbxkeabtdusyhrqklhaqreupakxkfzxgawqfwsaboszvlshwzhosojjotgyagygguzntrouhiweuomqptfjjqsxlbylhwtpssdlltgubczxslqjgxuqnmpynnlwjgmebrpokxjnbiltvbebyytnnjlcwyzignmhedwqbfdepqakrelrdfesqrumptwwgifmmbepiktxavhuavlfaqxqhreznbvvlakzeoomykkzftthoemqwliednfsqcnbexbimrvkdhllcesrlhhjsspvfupxwdybablotibypmjutclgjurbmhztboqatrdwsomnxnmocvixxvfiqwmednahdqhxjkvcyhpxxdmzzuyyqdjibvmfkmonfxmohhshpkhmntnoplphqyprveyfsmsxjfosmicdsjrieeytpnbhlsziwxnpmgoxneqbnufhfwrjbqcsdfarybzwaplmxckkgclvwqdbpumsmqkswmjwnkuqbicykoisqwoootrdpdvcuiuswfqmrkctsgrevcxnyncmivsxbpbxzxpwchiwtkroqisnmrbmefbmatmdknaklpgpyqlsccgunaibsloyqpnsibwuowebomrmcegejozypjzjunjmeygozcjqbnrpakdermjcckartbcppmbtkhkmmtcngteigjnxxyzaibtdcwutkvpwezisskfaeljmxyjwykwglqlnofhycwuivdbnpintuyhtyqpwaoelgpbuwiuyeqhbvkqlsfgmeoheexbhnhutxvnvfjwlzfmvpcghiowocdsjcvqrdmkcizxnivbianfpsnzabxqecinhgfyjrjlbikrrgsbgfgyxtzzwwpayapfgueroncpxogouyrdgzdfucfrywtywjeefkvtzxlw"

	// resp := minWindows(s, t)
	resp := minimalis(s, t)
	// resp := minimabis(s, t)

	if resp != jawaban {
		fmt.Println("\nSALAH!!!")
		fmt.Printf("resp: %v\n", resp)
	} else {
		fmt.Printf("resp: %v\n", resp)
	}
}

func minimalis(s string, t string) string {
	if len(t) > len(s) || t == "" {
		return ""
	}

	targetWindow := make(map[string]int, 0)

	for c := range t {
		key := string(t[c])
		targetWindow[key]++
	}

	n := len(s)
	k := len(targetWindow)

	left := 0
	right := 0
	count := 0
	subStrWin := make(map[string]int, k)
	listWindow := make([]string, 0)

	for right < n {
		currChar := string(s[right])
		if _, ok := targetWindow[currChar]; ok {
			subStrWin[currChar]++

			if subStrWin[currChar] == targetWindow[currChar] {
				count++
			}
		}

		for count == k {
			minWindow := s[left : right+1]
			listWindow = append(listWindow, minWindow)

			currHave := string(s[left])
			if _, ok := subStrWin[currHave]; ok {
				subStrWin[currHave]--
				if subStrWin[currHave] < targetWindow[currHave] {
					count--
				}
			}
			left++
		}

		right++
	}

	var (
		final     string
		minLength int = math.MinInt
	)
	for i := range listWindow {
		if len(listWindow[i]) <= minLength {
			minLength = len(listWindow[i])
			final = listWindow[i]
		}
	}

	return final

}
