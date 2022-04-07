//line date.rl:1
// Code generated by Ragel. DO NOT EDIT.

package date

// TODO: support more date formats
// than simply rfc3339

//line date.rl:69

func parse(data []byte) (year, month, day, hour, minute, second, nsec int, ok bool) {
	cs, p, pe, eof := 0, 0, len(data), -1
	// fractional component; divided by ten
	// for each decimal place after '.' that we scan
	frac, fracdig := 0, int(1e9)

//line date.go:21
	const date_start int = 1
	const date_first_final int = 32
	const date_error int = 0

	const date_en_main int = 1

//line date.go:29
	{
		cs = date_start
	}

//line date.go:34
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 27:
			goto st_case_27
		case 43:
			goto st_case_43
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		}
		goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 32 {
			goto st1
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st2
			}
		case data[p] >= 9:
			goto st1
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if 48 <= data[p] && data[p] <= 57 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if 48 <= data[p] && data[p] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 45 {
			goto tr6
		}
		goto st0
	tr6:
//line date.rl:13
		{
			year = int(data[p-4]-'0') * 1000
			year += int(data[p-3]-'0') * 100
			year += int(data[p-2]-'0') * 10
			year += int(data[p-1] - '0')
		}
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line date.go:201
		switch data[p] {
		case 48:
			goto st7
		case 49:
			goto st31
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if 49 <= data[p] && data[p] <= 57 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 45 {
			goto tr10
		}
		goto st0
	tr10:
//line date.rl:25
		{
			month = int(data[p-2]-'0') * 10
			month += int(data[p-1] - '0')
		}
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line date.go:239
		switch data[p] {
		case 48:
			goto st10
		case 51:
			goto st30
		}
		if 49 <= data[p] && data[p] <= 50 {
			goto st29
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 49 <= data[p] && data[p] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 32:
			goto tr15
		case 84:
			goto tr15
		}
		goto st0
	tr15:
//line date.rl:20
		{
			day = int(data[p-2]-'0') * 10
			day += int(data[p-1] - '0')
		}
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line date.go:283
		if data[p] == 50 {
			goto st28
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if 48 <= data[p] && data[p] <= 57 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 58 {
			goto tr19
		}
		goto st0
	tr19:
//line date.rl:30
		{
			hour = int(data[p-2]-'0') * 10
			hour += int(data[p-1] - '0')
		}
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line date.go:321
		if 48 <= data[p] && data[p] <= 53 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[p] && data[p] <= 57 {
			goto st17
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 58 {
			goto tr22
		}
		goto st0
	tr22:
//line date.rl:35
		{
			minute = int(data[p-2]-'0') * 10
			minute += int(data[p-1] - '0')
		}
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line date.go:356
		if data[p] == 54 {
			goto st27
		}
		if 48 <= data[p] && data[p] <= 53 {
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr25
		}
		goto st0
	tr25:
//line date.rl:39
		{
			second = int(data[p-1]-'0') * 10
			second += int(data[p] - '0')
		}
		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line date.go:385
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 46:
			goto st26
		case 90:
			goto st33
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st33
		}
		goto st0
	tr31:
//line date.rl:50
		{
			hoff := int(data[p-4]-'0')*10 + int(data[p-3]-'0')
			moff := int(data[p-1]-'0')*10 + int(data[p]-'0')
			if data[p-5] == '-' {
				hoff, moff = -hoff, -moff
			}
			hour, minute = hour-hoff, minute-moff
		}
		goto st33
	tr45:
//line date.rl:42
		second = 60
		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line date.go:422
		if data[p] == 32 {
			goto st33
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st33
		}
		goto st0
	tr46:
//line date.rl:42
		second = 60
		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line date.go:439
		if data[p] == 50 {
			goto st25
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 58 {
			goto st23
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if 48 <= data[p] && data[p] <= 53 {
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr31
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if 48 <= data[p] && data[p] <= 51 {
			goto st22
		}
		goto st0
	tr47:
//line date.rl:42
		second = 60
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line date.go:501
		if 48 <= data[p] && data[p] <= 57 {
			goto tr32
		}
		goto st0
	tr32:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line date.go:519
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr37
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr37:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line date.go:552
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr38
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr38:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line date.go:585
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr39
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr39:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line date.go:618
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr40
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr40:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line date.go:651
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr41
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr41:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line date.go:684
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr42
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr42:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line date.go:717
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr43
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr43:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line date.go:750
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr44
			}
		case data[p] >= 9:
			goto st33
		}
		goto st0
	tr44:
//line date.rl:44
		{
			frac *= 10
			frac += int(data[p] - '0')
			fracdig /= 10
		}
		goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line date.go:783
		switch data[p] {
		case 32:
			goto st33
		case 43:
			goto st20
		case 45:
			goto st20
		case 90:
			goto st33
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st33
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 48 {
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 32:
			goto tr45
		case 43:
			goto tr46
		case 45:
			goto tr46
		case 46:
			goto tr47
		case 90:
			goto tr45
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr45
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if 48 <= data[p] && data[p] <= 51 {
			goto st14
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if 48 <= data[p] && data[p] <= 57 {
			goto st11
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if 48 <= data[p] && data[p] <= 49 {
			goto st11
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if 48 <= data[p] && data[p] <= 50 {
			goto st8
		}
		goto st0
	st_out:
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 43:
//line date.rl:42
				second = 60
//line date.go:915
			}
		}

	_out:
		{
		}
	}

//line date.rl:81

	if cs < date_first_final {
		return
	}
	if frac != 0 {
		nsec = frac * fracdig
	}
	ok = true
	return
}
