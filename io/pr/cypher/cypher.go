package cypher

import (
	"errors"
	"math"
)

type (
	Rijndael interface {
		Encrypt(source []byte) (result []byte, err error)
		Decrypt(cipher []byte) (result []byte, err error)
	}
	rijndael struct {
		padding   Padding
		key       []byte
		iv        []byte
		blockSize byte
		rounds    byte
		ke        [][]uint32
		kd        [][]uint32
	}
	Padding interface {
		Encode(source []byte) (result []byte)
		Decode(source []byte) (result []byte)
	}
	padding struct {
		blockSize int
	}
)

func NewRijndael() Rijndael {
	return &rijndael{
		key:       []byte{97, 9, 7, 0, 185, 195, 184, 185, 66, 21, 153, 231, 156, 165, 123, 135, 190, 169, 50, 211, 121, 123, 173, 118, 99, 237, 77, 222, 10, 148, 60, 197},
		iv:        []byte{2, 36, 249, 119, 31, 219, 110, 14, 59, 213, 8, 215, 183, 149, 191, 46, 12, 189, 23, 105, 66, 104, 10, 99, 123, 18, 188, 98, 115, 219, 46, 187},
		blockSize: 32,
		rounds:    num_rounds[32][32],
		ke: [][]uint32{
			{1627981568, 3116611769, 1108711911, 2628090759, 3198759635, 2038148470, 1676496350, 177487045},
			{1122148711, 4213250526, 3107225657, 630324158, 2166898045, 4166225931, 2612977109, 2435506448},
			{2769972198, 1580954168, 3876581889, 3265137087, 2753772661, 1550888062, 3352195499, 1457819835},
			{3363471703, 2520904559, 1900937582, 3016755409, 3377663051, 2502139957, 1391317406, 67715365},
			{3254369957, 1471696330, 653675684, 2502313077, 3819182038, 1988605923, 610952829, 543369048},
			{708170770, 2106431960, 1534842236, 3462238473, 1760073431, 510599476, 973424457, 442916881},
			{971331248, 1147719528, 521345556, 3511585565, 1455926131, 1219004999, 1923992846, 1758178591},
			{1718654709, 572021149, 1024148361, 3963996308, 2561229649, 3489777942, 2729182232, 3395791111},
			{1671648129, 1102776860, 2091927957, 2432008449, 4165947181, 676234811, 2330100259, 1082469156},
			{4015154568, 2934534036, 3529057793, 1118611200, 3571323726, 4237333877, 1987188566, 922172530},
			{2981701005, 525355545, 3439946776, 2410007320, 2815954147, 1531484566, 758812352, 466465458},
			{1619923490, 2145278011, 3000448035, 1030946619, 2157451777, 3687875479, 4142533975, 3978805221},
			{1954990967, 190354252, 3113199471, 2231195732, 3754717217, 69006262, 4076273377, 533807364},
			{3775216055, 3932007163, 1406274964, 3610203584, 3520695451, 3586490157, 657686988, 954384584},
			{884986288, 3739384651, 2368763615, 1511990047, 1863952987, 3135034742, 2649707706, 2769130610}},
		kd: [][]uint32{
			{884986288, 3739384651, 2368763615, 1511990047, 1863952987, 3135034742, 2649707706, 2769130610},
			{1754555634, 328303840, 4054622170, 711094871, 1233595131, 256057647, 104469111, 1108739459},
			{1241344656, 2063982610, 3795655482, 3687771533, 391265339, 1187253204, 158937944, 1143740404},
			{6280532, 855149186, 2570648360, 972289719, 4123945297, 1368799215, 1337800844, 1297425580},
			{278781938, 849829846, 2881515946, 2697661855, 3220910075, 2757254846, 506138467, 48765984},
			{1990155493, 574259236, 2573676156, 185302069, 686834126, 463696197, 3128119773, 482546499},
			{3125980756, 1420156097, 3143478872, 2456572489, 1090209320, 861089931, 2714771608, 2796569246},
			{1467305242, 4009194133, 4026036889, 691141649, 2609579573, 1940420259, 2458062867, 123757062},
			{3425097479, 3112320911, 17760268, 3335097992, 2177822440, 3894636694, 3777717936, 2514711061},
			{2319179189, 1973741704, 3096264579, 3351678596, 897274589, 1777180286, 151524902, 1959269541},
			{1983663910, 4288623933, 3442071307, 2135666951, 1255188917, 1553386659, 1625655384, 2109735555},
			{3779748015, 2309192219, 850846262, 2992793100, 2967655515, 373712150, 1014187259, 488999643},
			{1822213636, 1760107188, 3138754605, 2161411130, 62528619, 2795811149, 708125165, 559336992},
			{3275615348, 74833072, 3556541081, 1002511383, 4004354865, 2770251046, 2358349984, 191098829},
			{1627981568, 3116611769, 1108711911, 2628090759, 3198759635, 2038148470, 1676496350, 177487045}},
		padding: NewPadding(),
	}
}

func (r rijndael) Encrypt(source []byte) (result []byte, err error) {
	var (
		ppt    = r.padding.Encode(source)
		offset = 0
		v      = make([]byte, len(r.iv))
	)
	copy(v, r.iv)
	for offset < len(ppt) {
		block := ppt[offset : offset+int(r.blockSize)]
		block = r.xOrBlock(block, v)
		block, err = r.encrypt(block)
		if err != nil {
			result = nil
			break
		}
		result = append(result, block...)
		offset += int(r.blockSize)
		v = block
	}
	return
}

func (r rijndael) Decrypt(cipher []byte) (result []byte, err error) {
	var (
		ppt       []byte
		offset    int
		v         = make([]byte, len(r.iv))
		decrypted []byte
	)
	copy(v, r.iv)
	for offset < len(cipher) {
		block := cipher[offset : offset+int(r.blockSize)]
		if decrypted, err = r.decrypt(block); err != nil {
			return nil, err
		}
		ppt = append(ppt, r.xOrBlock(decrypted, v)...)
		offset += int(r.blockSize)
		v = block
	}
	return r.padding.Decode(ppt), nil
}

func (r rijndael) xOrBlock(b1, b2 []byte) (result []byte) {
	for i := 0; i < int(r.blockSize); i++ {
		result = append(result, b1[i]^b2[i])
	}
	return
}

func (r rijndael) encrypt(source []byte) (result []byte, err error) {
	if len(source)%int(r.blockSize) != 0 {
		err = errors.New("wrong block length")
		return
	}
	var (
		rounds = len(r.ke) - 1
		sc     = 2
	)
	if r.blockSize == 4 {
		sc = 0
	} else if r.blockSize == 6 {
		sc = 1
	}
	bc := int(math.Floor(float64(r.blockSize) / 4.0))
	s1 := shifts[sc][1][0]
	s2 := shifts[sc][2][0]
	s3 := shifts[sc][3][0]
	a := make([]uint32, bc)
	var t []uint32
	for i := 0; i < bc; i++ {
		t = append(t,
			uint32(source[i*4])<<24|
				uint32(source[i*4+1])<<16|
				uint32(source[i*4+2])<<8|
				uint32(source[i*4+3])^r.ke[0][i])
	}
	for j := 1; j <= rounds; j++ {
		for i := 0; i < bc; i++ {
			a[i] =
				T1[(t[i]>>24)&0xFF] ^
					T2[(t[(i+int(s1))%bc]>>16)&0xFF] ^
					T3[(t[(i+int(s2))%bc]>>8)&0xFF] ^
					T4[t[(i+int(s3))%bc]&0xFF] ^ r.ke[j][i]
		}
		copy(t, a)
	}
	for i := 0; i < bc; i++ {
		tt := r.ke[rounds][i]

		j := uint32(S[int(byte(t[i]>>24)&0xFF)]) ^ (tt >> 24)
		result = append(result, byte(j)&0xFF)

		j = uint32(S[int(byte(t[(i+int(s1))%bc]>>16)&0xFF)]) ^ (tt >> 16)
		result = append(result, byte(j)&0xFF)

		j = uint32(S[int(byte(t[(i+int(s2))%bc]>>8)&0xFF)]) ^ (tt >> 8)
		result = append(result, byte(j)&0xFF)

		j = uint32(S[int(byte(t[(i+int(s3))%bc]&0xFF))]) ^ tt
		result = append(result, byte(j)&0xFF)
	}
	return result, nil
}

func (r rijndael) decrypt(cipher []byte) (result []byte, err error) {
	rounds := len(r.kd) - 1
	sc := 2
	if r.blockSize == 4 {
		sc = 0
	} else if r.blockSize == 6 {
		sc = 1
	}
	bc := int(math.Floor(float64(r.blockSize) / 4.0))
	s1 := shifts[sc][1][1]
	s2 := shifts[sc][2][1]
	s3 := shifts[sc][3][1]
	a := make([]uint32, bc)
	t := make([]uint32, bc)
	for i := 0; i < bc; i++ {
		t[i] = uint32(cipher[i*4])<<24 |
			uint32(cipher[i*4+1])<<16 |
			uint32(cipher[i*4+2])<<8 |
			uint32(cipher[i*4+3]) ^ r.kd[0][i]
	}
	for j := 1; j < rounds; j++ {
		for i := 0; i < bc; i++ {
			a[i] =
				T5[(t[i]>>24)&0xFF] ^
					T6[(t[(i+int(s1))%bc]>>16)&0xFF] ^
					T7[(t[(i+int(s2))%bc]>>8)&0xFF] ^
					T8[t[(i+int(s3))%bc]&0xFF] ^ r.kd[j][i]
		}
		copy(t, a)
	}
	for i := 0; i < bc; i++ {
		tt := r.kd[rounds][i]
		result = append(result, Si[int(byte(t[i]>>24)&0xFF)]^byte(tt>>24))
		result = append(result, Si[int(byte(t[(i+int(s1))%bc]>>16)&0xFF)]^byte(tt>>16))
		result = append(result, Si[int(byte(t[(i+int(s2))%bc]>>8)&0xFF)]^byte(tt>>8))
		result = append(result, Si[int(byte(t[(i+int(s3))%bc]&0xFF))]^byte(tt))
	}
	return
}

func NewPadding() Padding {
	return &padding{blockSize: 32}
}

func (p padding) Encode(source []byte) []byte {
	count := p.blockSize - ((len(source)+p.blockSize-1)%p.blockSize + 1)
	for i := 0; i < count; i++ {
		source = append(source, 0)
	}
	return source
}

func (p padding) Decode(source []byte) []byte {
	offset := len(source)
	if offset == 0 {
		return []byte{}
	}
	end := offset - p.blockSize + 1
	for offset > end {
		offset--
		if source[offset] != 0 {
			return source[:offset+1]
		}
	}
	return source[:end]
}
