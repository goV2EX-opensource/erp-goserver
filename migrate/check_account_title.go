package migrate

import (
	"errors"
	"v2ex/go-erp/dbutil"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//DicAccountTitle : 会计科目词典表
type DicAccountTitle struct {
	ID   uint32 `gorm:"type:INT UNSIGNED NOT NULL AUTO_INCREMENT"`
	Code uint16 `gorm:"unique;type:MEDIUMINT UNSIGNED NOT NULL"`
	Name string `gorm:"type:char(24);not null"`
	Memo string `gorm:"type:char(24);not null"`
}

//CheckATD 确认会计科目词典存在
func CheckATD() {
	log.Info("Checking ATD table")
	db := dbutil.GetDB()
	if !db.Migrator().HasTable(&DicAccountTitle{}) {
		log.Info("NO ATD TABLE. PREPARE INIT ONE.")
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&DicAccountTitle{})
		log.Info("ATD table inited")
	}

	var atdic DicAccountTitle

	if result := db.First(&atdic); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Info("NO ATD DATA RECORED, PREPARE INIT")

			tx := db.Begin()

			var datObj DicAccountTitle

			datObj = DicAccountTitle{Code: 1001, Name: "库存现金", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1002, Name: "银行存款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1003, Name: "存放中央银行款项", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1011, Name: "存放同业", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1015, Name: "其它货币基金", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1021, Name: "结算备付金", Memo: "证券专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1031, Name: "存出保证金", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1051, Name: "拆出资金", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1101, Name: "交易性金融资产", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1111, Name: "买入返售金融资产", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1121, Name: "应收票据", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1122, Name: "应收帐款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1123, Name: "预付帐款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1131, Name: "应收股利", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1132, Name: "应收利息", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1211, Name: "应收保护储金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1221, Name: "应收代位追偿款", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1222, Name: "应收分保帐款", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1223, Name: "应收分保未到期责任准备金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1224, Name: "应收分保保险责任准备金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1231, Name: "其它应收款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1241, Name: "坏帐准备", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1251, Name: "贴现资产", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1301, Name: "贷款", Memo: "银行和保险共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1302, Name: "贷款损失准备", Memo: "银行和保险共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1311, Name: "代理兑付证券", Memo: "银行和保险共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1321, Name: "代理业务资产", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1401, Name: "材料采购", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1402, Name: "在途物资", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1403, Name: "原材料", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1404, Name: "材料成本差异", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1406, Name: "库存商品", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1407, Name: "发出商品", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1410, Name: "商品进销差价", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1411, Name: "委托加工物资", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1412, Name: "包装物及低值易耗品", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1421, Name: "消耗性物物资产", Memo: "农业专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1431, Name: "周转材料", Memo: "建造承包商专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1441, Name: "贵金属", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1442, Name: "抵债资产", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1451, Name: "损余物资", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1461, Name: "存货跌价准备", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1501, Name: "待摊费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1511, Name: "独立帐户资产", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1521, Name: "持有至到期投资", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1522, Name: "持有至到期投资减值准备", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1523, Name: "可供出售金融资产", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1524, Name: "长期股权投资", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1525, Name: "长期股权投资减值准备", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1526, Name: "投资性房地产", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1531, Name: "长期应收款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1541, Name: "未实现融资收益", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1551, Name: "存出资本保证金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1601, Name: "固定资产", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1602, Name: "累计折旧", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1603, Name: "固定资产减值准备", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1604, Name: "在建工程", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1605, Name: "工程物资", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1606, Name: "固定资产清理", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1611, Name: "融资租赁资产", Memo: "租赁专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1612, Name: "未担保余值", Memo: "租赁专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1621, Name: "生产性生物资产", Memo: "农业专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1622, Name: "生产性生物资产累计折旧", Memo: "农业专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1623, Name: "公益性生物资产", Memo: "农业专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1631, Name: "油气资产", Memo: "石油天然气开采专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1632, Name: "累计折耗", Memo: "石油天然气开采专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1701, Name: "无形资产", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1702, Name: "累计摊销", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1703, Name: "无形资产减值准备", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1711, Name: "商誉", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1801, Name: "长期待摊费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1811, Name: "递延所得资产", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 1901, Name: "待处理财产损益", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2001, Name: "短期借款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2002, Name: "存入保证金", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2003, Name: "拆入资金", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2004, Name: "向中央银行借款", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2011, Name: "同业存放", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2012, Name: "吸收存款", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2021, Name: "贴现负债", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2101, Name: "交易性金融负债", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2111, Name: "专出回购金融资产款", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2201, Name: "应付票据", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2202, Name: "应付帐款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2205, Name: "预收帐款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2211, Name: "应付职工薪酬", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2221, Name: "应交税费", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2231, Name: "应付股利", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2232, Name: "应付利息", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2241, Name: "其他应付款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2251, Name: "应付保户红利", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2261, Name: "应付分保帐款", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2311, Name: "代理买卖证券款", Memo: "证券专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2312, Name: "代理承销证券款", Memo: "证券和银行共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2313, Name: "代理兑付证券款", Memo: "证券和银行共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2314, Name: "代理业务负债", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2401, Name: "预提费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2411, Name: "预计负债", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2501, Name: "递延收益", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2601, Name: "长期借款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2602, Name: "长期债券", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2701, Name: "未到期责任准备金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2702, Name: "保险责任准备金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2711, Name: "保户储金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2721, Name: "独立帐户负债", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2801, Name: "长期应付款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2802, Name: "未确认融资费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2811, Name: "专项应付款", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 2901, Name: "递延所得税负债", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 3001, Name: "清算资金往来", Memo: "银行专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 3002, Name: "外汇买卖", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 3101, Name: "衍生工具", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 3201, Name: "套期工具", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 3202, Name: "被套期项目", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 4001, Name: "实收资本", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 4002, Name: "资本公积", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 4101, Name: "盈余公积", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 4102, Name: "一般风险准备", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 4103, Name: "本年利润", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 4104, Name: "利润分配", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 4201, Name: "库存股", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 5001, Name: "生产成本", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 5101, Name: "制造费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 5201, Name: "劳务成本", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 5301, Name: "研发支出", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 5401, Name: "工程施工", Memo: "建造承包商专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 5402, Name: "工程结算", Memo: "建造承包商专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 5403, Name: "机械作业", Memo: "建造承包商专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6001, Name: "主营业务收入", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6011, Name: "利息收入", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6021, Name: "手续费收入", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6031, Name: "保费收入", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6032, Name: "分保费收入", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6041, Name: "租赁收入", Memo: "租赁专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6051, Name: "其他业务收入", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6061, Name: "汇兑损益", Memo: "金融专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6101, Name: "公允价值变动损益", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6111, Name: "投资收益", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6201, Name: "摊回保险责任准备金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6202, Name: "摊回赔付支出", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6203, Name: "摊回分保费用", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6301, Name: "营业外收入", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6401, Name: "主营业务成本", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6402, Name: "其它业务成本", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6405, Name: "营业税金及附加", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6411, Name: "利息支出", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6421, Name: "手续费支出", Memo: "金融共用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6501, Name: "提取未到期责任准备金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6502, Name: "撮保险责任准备金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6511, Name: "赔付支出", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6521, Name: "保户红利支出", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6531, Name: "退保金", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6541, Name: "分出保费", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6542, Name: "分保费用", Memo: "保险专用"}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6601, Name: "销售费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6602, Name: "管理费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6603, Name: "财务费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6604, Name: "勘探费用", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6701, Name: "资产减值损失", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6711, Name: "营业外支出", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6801, Name: "所得税", Memo: ""}
			_ = tx.Create(&datObj)
			datObj = DicAccountTitle{Code: 6901, Name: "以前年度损益调整", Memo: ""}
			_ = tx.Create(&datObj)

			if tx.Error != nil {
				tx.Rollback()
				log.Fatal("DB FATAL ERROR")
			} else {
				tx.Commit()
			}
		} else {
			log.Fatal("DB FATAL ERROR")
		}
	}
}
