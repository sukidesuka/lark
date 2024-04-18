package card

// Color 颜色
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/enumerations-for-fields-related-to-color
type Color = string

const (
	ColorDefault      Color = "default"
	ColorBgWhite      Color = "bg-white"      // 浅色主题:  N00, 深色主题:/#ffffff N50/#1A1A1A
	ColorWhite        Color = "white"         // 浅色主题:  /, 深色主题:   /
	ColorBlue         Color = "blue"          // 蓝, 浅色主题: #1456F0, 深色主题: #75A4FF
	ColorBlue50       Color = "blue-50"       // 蓝, 浅色主题: #F0F4FF, 深色主题: #152340
	ColorBlue100      Color = "blue-100"      // 蓝, 浅色主题: #E0E9FF, 深色主题: #173166
	ColorBlue200      Color = "blue-200"      // 蓝, 浅色主题: #C2D4FF, 深色主题: #194294
	ColorBlue300      Color = "blue-300"      // 蓝, 浅色主题: #94B4FF, 深色主题: #2655B6
	ColorBlue350      Color = "blue-350"      // 蓝, 浅色主题: #7AA2FF, 深色主题: #275FCE
	ColorBlue400      Color = "blue-400"      // 蓝, 浅色主题: #5083FB, 深色主题: #3370EB
	ColorBlue500      Color = "blue-500"      // 蓝, 浅色主题: #336DF4, 深色主题: #4C88FF
	ColorBlue600      Color = "blue-600"      // 蓝, 浅色主题: #1456F0, 深色主题: #75A4FF
	ColorBlue700      Color = "blue-700"      // 蓝, 浅色主题: #0442D2, 深色主题: #8FB4FF
	ColorBlue800      Color = "blue-800"      // 蓝, 浅色主题: #002F9E, 深色主题: #BDD2FF
	ColorBlue900      Color = "blue-900"      // 蓝, 浅色主题: #002270, 深色主题: #E0E9FF
	ColorCarmine      Color = "carmine"       // 粉, 浅色主题: #B82879, 深色主题: #ED77BA
	ColorCarmine50    Color = "carmine-50"    // 粉, 浅色主题: #FEF0F8, 深色主题: #3A182B
	ColorCarmine100   Color = "carmine-100"   // 粉, 浅色主题: #FEE2F2, 深色主题: #591C3F
	ColorCarmine200   Color = "carmine-200"   // 粉, 浅色主题: #F8C4E1, 深色主题: #782B57
	ColorCarmine300   Color = "carmine-300"   // 粉, 浅色主题: #F598CC, 深色主题: #94386C
	ColorCarmine350   Color = "carmine-350"   // 粉, 浅色主题: #EB78B8, 深色主题: #AB417D
	ColorCarmine400   Color = "carmine-400"   // 粉, 浅色主题: #DF58A5, 深色主题: #C24A8E
	ColorCarmine500   Color = "carmine-500"   // 粉, 浅色主题: #CC398C, 深色主题: #DB5EA4
	ColorCarmine600   Color = "carmine-600"   // 粉, 浅色主题: #B82879, 深色主题: #ED77BA
	ColorCarmine700   Color = "carmine-700"   // 粉, 浅色主题: #9D1562, 深色主题: #FC94CF
	ColorCarmine800   Color = "carmine-800"   // 粉, 浅色主题: #730744, 深色主题: #FFC2E5
	ColorCarmine900   Color = "carmine-900"   // 粉, 浅色主题: #550C35, 深色主题: #FFE0F2
	ColorGreen        Color = "green"         // 绿, 浅色主题: #1A7526, 深色主题: #51BA43
	ColorGreen50      Color = "green-50"      // 绿, 浅色主题: #E4FAE1, 深色主题: #0E2B0A
	ColorGreen100     Color = "green-100"     // 绿, 浅色主题: #D0F5CE, 深色主题: #173B12
	ColorGreen200     Color = "green-200"     // 绿, 浅色主题: #95E599, 深色主题: #21511A
	ColorGreen300     Color = "green-300"     // 绿, 浅色主题: #5CD168, 深色主题: #296621
	ColorGreen350     Color = "green-350"     // 绿, 浅色主题: #35BD4B, 深色主题: #2F7526
	ColorGreen400     Color = "green-400"     // 绿, 浅色主题: #32A645, 深色主题: #35872A
	ColorGreen500     Color = "green-500"     // 绿, 浅色主题: #258832, 深色主题: #419E34
	ColorGreen600     Color = "green-600"     // 绿, 浅色主题: #1A7526, 深色主题: #51BA43
	ColorGreen700     Color = "green-700"     // 绿, 浅色主题: #0B6017, 深色主题: #69CC5C
	ColorGreen800     Color = "green-800"     // 绿, 浅色主题: #04430C, 深色主题: #99E490
	ColorGreen900     Color = "green-900"     // 绿, 浅色主题: #022C07, 深色主题: #CBF5C6
	ColorIndigo       Color = "indigo"        // 靛蓝, 浅色主题: #4752E6, 深色主题: #9499F7
	ColorIndigo50     Color = "indigo-50"     // 靛蓝, 浅色主题: #F2F3FD, 深色主题: #1E204A
	ColorIndigo100    Color = "indigo-100"    // 靛蓝, 浅色主题: #E9EAFB, 深色主题: #2A2D69
	ColorIndigo200    Color = "indigo-200"    // 靛蓝, 浅色主题: #CCCFF9, 深色主题: #373D90
	ColorIndigo300    Color = "indigo-300"    // 靛蓝, 浅色主题: #ABB0F2, 深色主题: #474FB8
	ColorIndigo350    Color = "indigo-350"    // 靛蓝, 浅色主题: #959BEE, 深色主题: #515AD6
	ColorIndigo400    Color = "indigo-400"    // 靛蓝, 浅色主题: #757DF0, 深色主题: #5E68E8
	ColorIndigo500    Color = "indigo-500"    // 靛蓝, 浅色主题: #5B65F5, 深色主题: #7B83F7
	ColorIndigo600    Color = "indigo-600"    // 靛蓝, 浅色主题: #4752E6, 深色主题: #9499F7
	ColorIndigo700    Color = "indigo-700"    // 靛蓝, 浅色主题: #333DCC, 深色主题: #AAAFF8
	ColorIndigo800    Color = "indigo-800"    // 靛蓝, 浅色主题: #1E27A4, 深色主题: #CDD0F9
	ColorIndigo900    Color = "indigo-900"    // 靛蓝, 浅色主题: #151B70, 深色主题: #E7E9FE
	ColorLime         Color = "lime"          // 柠檬绿, 浅色主题: #5C6D08, 深色主题: #93AF04
	ColorLime50       Color = "lime-50"       // 柠檬绿, 浅色主题: #F2F8D3, 深色主题: #212702
	ColorLime100      Color = "lime-100"      // 柠檬绿, 浅色主题: #E3F0A3, 深色主题: #303804
	ColorLime200      Color = "lime-200"      // 柠檬绿, 浅色主题: #C8DD5F, 深色主题: #404C06
	ColorLime300      Color = "lime-300"      // 柠檬绿, 浅色主题: #A2C10B, 深色主题: #53610A
	ColorLime350      Color = "lime-350"      // 柠檬绿, 浅色主题: #91AD00, 深色主题: #5F700A
	ColorLime400      Color = "lime-400"      // 柠檬绿, 浅色主题: #7B9207, 深色主题: #6B7F05
	ColorLime500      Color = "lime-500"      // 柠檬绿, 浅色主题: #6B7F06, 深色主题: #82990A
	ColorLime600      Color = "lime-600"      // 柠檬绿, 浅色主题: #5C6D08, 深色主题: #93AF04
	ColorLime700      Color = "lime-700"      // 柠檬绿, 浅色主题: #4A5804, 深色主题: #A6C313
	ColorLime800      Color = "lime-800"      // 柠檬绿, 浅色主题: #333D00, 深色主题: #C2E12D
	ColorLime900      Color = "lime-900"      // 柠檬绿, 浅色主题: #262E00, 深色主题: #E3F391
	ColorGrey         Color = "grey"          // 中性色-neutral, 浅色主题: #646a73 深色主题: #a6a6a6
	ColorGrey00       Color = "grey-00"       // 中性色-neutral, 浅色主题: #ffffff 深色主题: #0A0A0A
	ColorGrey50       Color = "grey-50"       // 中性色-neutral, 浅色主题: #f5f6f7 深色主题: #1A1A1A
	ColorGrey100      Color = "grey-100"      // 中性色-neutral, 浅色主题: #f2f3f5 深色主题: #292929
	ColorGrey200      Color = "grey-200"      // 中性色-neutral, 浅色主题: #eff0f1 深色主题: #373737
	ColorGrey300      Color = "grey-300"      // 中性色-neutral, 浅色主题: #dee0e3 深色主题: #434343
	ColorGrey350      Color = "grey-350"      // 中性色-neutral, 浅色主题: #d0d3d6 深色主题: #505050
	ColorGrey400      Color = "grey-400"      // 中性色-neutral, 浅色主题: #bbbfc4 深色主题: #5f5f5f
	ColorGrey500      Color = "grey-500"      // 中性色-neutral, 浅色主题: #8f959e 深色主题: #757575
	ColorGrey600      Color = "grey-600"      // 中性色-neutral, 浅色主题: #646a73 深色主题: #a6a6a6
	ColorGrey650      Color = "grey-650"      // 中性色-neutral, 浅色主题: #51565d 深色主题: #cfcfcf
	ColorGrey700      Color = "grey-700"      // 中性色-neutral, 浅色主题: #373c43 深色主题: #e0e0e0
	ColorGrey800      Color = "grey-800"      // 中性色-neutral, 浅色主题: #2b2f36 深色主题: #e8e8e8
	ColorGrey900      Color = "grey-900"      // 中性色-neutral, 浅色主题: #1f2329 深色主题: #ebebeb
	ColorGrey950      Color = "grey-950"      // 中性色-neutral, 浅色主题: #0f1114 深色主题: #f8f8f8
	ColorGrey1000     Color = "grey-1000"     // 中性色-neutral, 浅色主题: #000000 深色主题: #ffffff
	ColorOrange       Color = "orange"        // 橙, 浅色主题: #A44904, 深色主题: #F3871B
	ColorOrange50     Color = "orange-50"     // 橙, 浅色主题: #FFF3E5, 深色主题: #33210B
	ColorOrange100    Color = "orange-100"    // 橙, 浅色主题: #FEE7CD, 深色主题: #4A2B10
	ColorOrange200    Color = "orange-200"    // 橙, 浅色主题: #FEC48B, 深色主题: #683A12
	ColorOrange300    Color = "orange-300"    // 橙, 浅色主题: #FF9D4C, 深色主题: #8A4A19
	ColorOrange350    Color = "orange-350"    // 橙, 浅色主题: #FF811A, 深色主题: #A15317
	ColorOrange400    Color = "orange-400"    // 橙, 浅色主题: #ED6D0C, 深色主题: #B85E1A
	ColorOrange500    Color = "orange-500"    // 橙, 浅色主题: #C25705, 深色主题: #DB7018
	ColorOrange600    Color = "orange-600"    // 橙, 浅色主题: #A44904, 深色主题: #F3871B
	ColorOrange700    Color = "orange-700"    // 橙, 浅色主题: #853A05, 深色主题: #F89E44
	ColorOrange800    Color = "orange-800"    // 橙, 浅色主题: #642B02, 深色主题: #FEC88B
	ColorOrange900    Color = "orange-900"    // 橙, 浅色主题: #3B1A02, 深色主题: #FEE7CD
	ColorPurple       Color = "purple"        // 紫, 浅色主题: #7A35F0, 深色主题: #B88FFE
	ColorPurple50     Color = "purple-50"     // 紫, 浅色主题: #F5F0FE, 深色主题: #2B194A
	ColorPurple100    Color = "purple-100"    // 紫, 浅色主题: #EFE6FE, 深色主题: #3F2073
	ColorPurple200    Color = "purple-200"    // 紫, 浅色主题: #DCC9FD, 深色主题: #5529A3
	ColorPurple300    Color = "purple-300"    // 紫, 浅色主题: #C8A9FC, 深色主题: #6C39C6
	ColorPurple350    Color = "purple-350"    // 紫, 浅色主题: #B791FA, 深色主题: #7C4AD4
	ColorPurple400    Color = "purple-400"    // 紫, 浅色主题: #9F6FF1, 深色主题: #8C55EC
	ColorPurple500    Color = "purple-500"    // 紫, 浅色主题: #8D55ED, 深色主题: #A575FA
	ColorPurple600    Color = "purple-600"    // 紫, 浅色主题: #7A35F0, 深色主题: #B88FFE
	ColorPurple700    Color = "purple-700"    // 紫, 浅色主题: #611FD6, 深色主题: #C5A3FF
	ColorPurple800    Color = "purple-800"    // 紫, 浅色主题: #4811A6, 深色主题: #DBC8FD
	ColorPurple900    Color = "purple-900"    // 紫, 浅色主题: #2F0080, 深色主题: #EFE5FF
	ColorRed          Color = "red"           // 红, 浅色主题: #C02A26, 深色主题: #F6827E
	ColorRed50        Color = "red-50"        // 红, 浅色主题: #FEF0F0, 深色主题: #3D1A19
	ColorRed100       Color = "red-100"       // 红, 浅色主题: #FEE3E2, 深色主题: #591F1D
	ColorRed200       Color = "red-200"       // 红, 浅色主题: #FDC6C4, 深色主题: #7B2524
	ColorRed300       Color = "red-300"       // 红, 浅色主题: #F89E9B, 深色主题: #A03331
	ColorRed350       Color = "red-350"       // 红, 浅色主题: #FF7570, 深色主题: #B33A37
	ColorRed400       Color = "red-400"       // 红, 浅色主题: #F54A45, 深色主题: #D14642
	ColorRed500       Color = "red-500"       // 红, 浅色主题: #E22E28, 深色主题: #F05B56
	ColorRed600       Color = "red-600"       // 红, 浅色主题: #C02A26, 深色主题: #F6827E
	ColorRed700       Color = "red-700"       // 红, 浅色主题: #A11C17, 深色主题: #F89E9B
	ColorRed800       Color = "red-800"       // 红, 浅色主题: #741915, 深色主题: #FDC6C4
	ColorRed900       Color = "red-900"       // 红, 浅色主题: #590603, 深色主题: #FEE3E2
	ColorSunflower    Color = "sunflower"     // 向日葵黄, 浅色主题: #8F7C00, 深色主题: #F5DF36
	ColorSunflower50  Color = "sunflower-50"  // 向日葵黄, 浅色主题: #FFFFDB, 深色主题: #29250A
	ColorSunflower100 Color = "sunflower-100" // 向日葵黄, 浅色主题: #FFFCA3, 深色主题: #38320C
	ColorSunflower200 Color = "sunflower-200" // 向日葵黄, 浅色主题: #FFF67A, 深色主题: #574D01
	ColorSunflower300 Color = "sunflower-300" // 向日葵黄, 浅色主题: #FFF258, 深色主题: #7A6A01
	ColorSunflower350 Color = "sunflower-350" // 向日葵黄, 浅色主题: #FFE928, 深色主题: #9C8702
	ColorSunflower400 Color = "sunflower-400" // 向日葵黄, 浅色主题: #E5CE00, 深色主题: #C9B218
	ColorSunflower500 Color = "sunflower-500" // 向日葵黄, 浅色主题: #C2A800, 深色主题: #E5CD17
	ColorSunflower600 Color = "sunflower-600" // 向日葵黄, 浅色主题: #8F7C00, 深色主题: #F5DF36
	ColorSunflower700 Color = "sunflower-700" // 向日葵黄, 浅色主题: #5C4F00, 深色主题: #F7E663
	ColorSunflower800 Color = "sunflower-800" // 向日葵黄, 浅色主题: #423700, 深色主题: #FAED7A
	ColorSunflower900 Color = "sunflower-900" // 向日葵黄, 浅色主题: #2C2502, 深色主题: #FFF59E
	ColorTurquoise    Color = "turquoise"     // 青, 浅色主题: #067062, 深色主题: #1AB7A1
	ColorTurquoise50  Color = "turquoise-50"  // 青, 浅色主题: #E2F8F5, 深色主题: #132926
	ColorTurquoise100 Color = "turquoise-100" // 青, 浅色主题: #C4F2EC, 深色主题: #173B36
	ColorTurquoise200 Color = "turquoise-200" // 青, 浅色主题: #6FE8D8, 深色主题: #1D4E47
	ColorTurquoise300 Color = "turquoise-300" // 青, 浅色主题: #33D6C0, 深色主题: #25665E
	ColorTurquoise350 Color = "turquoise-350" // 青, 浅色主题: #2DBEAB, 深色主题: #227369
	ColorTurquoise400 Color = "turquoise-400" // 青, 浅色主题: #10A893, 深色主题: #198578
	ColorTurquoise500 Color = "turquoise-500" // 青, 浅色主题: #0F8575, 深色主题: #1FA18F
	ColorTurquoise600 Color = "turquoise-600" // 青, 浅色主题: #067062, 深色主题: #1AB7A1
	ColorTurquoise700 Color = "turquoise-700" // 青, 浅色主题: #045D51, 深色主题: #17CFB5
	ColorTurquoise800 Color = "turquoise-800" // 青, 浅色主题: #03443B, 深色主题: #65E7D5
	ColorTurquoise900 Color = "turquoise-900" // 青, 浅色主题: #02312A, 深色主题: #B7F7EF
	ColorViolet       Color = "violet"        // 紫红, 浅色主题: #A630A6, 深色主题: #E17FE1
	ColorViolet50     Color = "violet-50"     // 紫红, 浅色主题: #FCEEFC, 深色主题: #3B153B
	ColorViolet100    Color = "violet-100"    // 紫红, 浅色主题: #F9E2F9, 深色主题: #541854
	ColorViolet200    Color = "violet-200"    // 紫红, 浅色主题: #F3C4F3, 深色主题: #712871
	ColorViolet300    Color = "violet-300"    // 紫红, 浅色主题: #E59CE5, 深色主题: #8B378B
	ColorViolet350    Color = "violet-350"    // 紫红, 浅色主题: #DE81DE, 深色主题: #A43DA4
	ColorViolet400    Color = "violet-400"    // 紫红, 浅色主题: #CF5ECF, 深色主题: #B54AB5
	ColorViolet500    Color = "violet-500"    // 紫红, 浅色主题: #BF3DBF, 深色主题: #D661D6
	ColorViolet600    Color = "violet-600"    // 紫红, 浅色主题: #A630A6, 深色主题: #E17FE1
	ColorViolet700    Color = "violet-700"    // 紫红, 浅色主题: #872787, 深色主题: #E99BE9
	ColorViolet800    Color = "violet-800"    // 紫红, 浅色主题: #6A116A, 深色主题: #F4C3F4
	ColorViolet900    Color = "violet-900"    // 紫红, 浅色主题: #520052, 深色主题: #FCDFFC
	ColorWathet       Color = "wathet"        // 天蓝, 浅色主题: #076A94, 深色主题: #25B2E5
	ColorWathet50     Color = "wathet-50"     // 天蓝, 浅色主题: #E7F8FE, 深色主题: #152830
	ColorWathet100    Color = "wathet-100"    // 天蓝, 浅色主题: #CAEFFC, 深色主题: #103647
	ColorWathet200    Color = "wathet-200"    // 天蓝, 浅色主题: #97DCFC, 深色主题: #164359
	ColorWathet300    Color = "wathet-300"    // 天蓝, 浅色主题: #3EC3F7, 深色主题: #135A78
	ColorWathet350    Color = "wathet-350"    // 天蓝, 浅色主题: #25B0E7, 深色主题: #146C91
	ColorWathet400    Color = "wathet-400"    // 天蓝, 浅色主题: #1295CA, 深色主题: #1A7FAB
	ColorWathet500    Color = "wathet-500"    // 天蓝, 浅色主题: #047FB0, 深色主题: #1099CC
	ColorWathet600    Color = "wathet-600"    // 天蓝, 浅色主题: #076A94, 深色主题: #25B2E5
	ColorWathet700    Color = "wathet-700"    // 天蓝, 浅色主题: #0F587A, 深色主题: #51C3F0
	ColorWathet800    Color = "wathet-800"    // 天蓝, 浅色主题: #06415C, 深色主题: #89DFFE
	ColorWathet900    Color = "wathet-900"    // 天蓝, 浅色主题: #072B3D, 深色主题: #C7F0FF
	ColorYellow       Color = "yellow"        // 黄, 浅色主题: #865B03, 深色主题: #FBCB46
	ColorYellow50     Color = "yellow-50"     // 黄, 浅色主题: #FBF4DF, 深色主题: #30250A
	ColorYellow100    Color = "yellow-100"    // 黄, 浅色主题: #FAEDC2, 深色主题: #473409
	ColorYellow200    Color = "yellow-200"    // 黄, 浅色主题: #FCDF7E, 深色主题: #63470F
	ColorYellow300    Color = "yellow-300"    // 黄, 浅色主题: #FAD355, 深色主题: #8A6419
	ColorYellow350    Color = "yellow-350"    // 黄, 浅色主题: #FFC60A, 深色主题: #AD7D15
	ColorYellow400    Color = "yellow-400"    // 黄, 浅色主题: #D99904, 深色主题: #D49B0B
	ColorYellow500    Color = "yellow-500"    // 黄, 浅色主题: #AD7A03, 深色主题: #F0B622
	ColorYellow600    Color = "yellow-600"    // 黄, 浅色主题: #865B03, 深色主题: #FBCB46
	ColorYellow700    Color = "yellow-700"    // 黄, 浅色主题: #6F4A01, 深色主题: #FCD456
	ColorYellow800    Color = "yellow-800"    // 黄, 浅色主题: #573601, 深色主题: #FFDE75
	ColorYellow900    Color = "yellow-900"    // 黄, 浅色主题: #382201, 深色主题: #FFEAA3
)