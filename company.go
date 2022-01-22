package creditcard

// Company holds a short and long names of who has issued the credit card
type Company struct {
	Long  string `json:"long,omitempty"`
	Short string `json:"short,omitempty"`
}

type CompanyName string

const (
	CompanyAmericanExpress         CompanyName = "American Express"
	CompanyAura                    CompanyName = "Aura"
	CompanyBankCard                CompanyName = "Bankcard"
	CompanyCabal                   CompanyName = "Cabal"
	CompanyChinaUnionPay           CompanyName = "China UnionPay"
	CompanyDankort                 CompanyName = "Dankort"
	CompanyDinersClubCarteBlance   CompanyName = "Diners Club Carte Blanche"
	CompanyDinersClubEnRoute       CompanyName = "Diners Club enRoute"
	CompanyDinersClubInternational CompanyName = "Diners Club International"
	CompanyDiscover                CompanyName = "Discover"
	CompanyElo                     CompanyName = "Elo"
	CompanyHipercard               CompanyName = "Hipercard"
	CompanyInstaPayment            CompanyName = "InstaPayment"
	CompanyInterPayment            CompanyName = "InterPayment"
	CompanyJCB                     CompanyName = "JCB"
	CompanyMaestro                 CompanyName = "Maestro"
	CompanyMasterCard              CompanyName = "MasterCard"
	CompanyNaranja                 CompanyName = "Naranja"
	CompanyVisa                    CompanyName = "Visa"
	CompanyVisaElectron            CompanyName = "Visa Electron"
)

const (
	CompanyNameLongAmericanExpress         = "American Express"
	CompanyNameLongAura                    = "Aura"
	CompanyNameLongBankcard                = "Bankcard"
	CompanyNameLongCabal                   = "Cabal"
	CompanyNameLongChinaUnionPay           = "China UnionPay"
	CompanyNameLongDankort                 = "Dankort"
	CompanyNameLongDinersClubCarteBlanche  = "Diners Club Carte Blanche"
	CompanyNameLongDinersClubEnRoute       = "Diners Club enRoute"
	CompanyNameLongDinersClubInternational = "Diners Club International"
	CompanyNameLongDiscover                = "Discover"
	CompanyNameLongElo                     = "Elo"
	CompanyNameLongHipercard               = "Hipercard"
	CompanyNameLongInstaPayment            = "InstaPayment"
	CompanyNameLongInterPayment            = "InterPayment"
	CompanyNameLongJCB                     = "JCB"
	CompanyNameLongMaestro                 = "Maestro"
	CompanyNameLongMasterCard              = "MasterCard"
	CompanyNameLongNaranja                 = "Naranja"
	CompanyNameLongVisa                    = "Visa"
	CompanyNameLongVisaElectron            = "Visa Electron"

	CompanyNameShortAmericanExpress         = "amex"
	CompanyNameShortAura                    = "aura"
	CompanyNameShortBankcard                = "bankcard"
	CompanyNameShortCabal                   = "cabal"
	CompanyNameShortChinaUnionPay           = "china unionpay"
	CompanyNameShortDankort                 = "dankort"
	CompanyNameShortDinersClubCarteBlanche  = "diners club carte blanche"
	CompanyNameShortDinersClubEnRoute       = "diners club enroute"
	CompanyNameShortDinersClubInternational = "diners club international"
	CompanyNameShortDiscover                = "discover"
	CompanyNameShortElo                     = "elo"
	CompanyNameShortHipercard               = "hipercard"
	CompanyNameShortInstaPayment            = "instapayment"
	CompanyNameShortInterPayment            = "interpayment"
	CompanyNameShortJCB                     = "jcb"
	CompanyNameShortMaestro                 = "maestro"
	CompanyNameShortMasterCard              = "mastercard"
	CompanyNameShortNaranja                 = "naranja"
	CompanyNameShortVisa                    = "visa"
	CompanyNameShortVisaElectron            = "visa electron"
)

var companies = map[CompanyName]Company{
	CompanyAmericanExpress: {
		Long:  CompanyNameLongAmericanExpress,
		Short: CompanyNameShortAmericanExpress,
	},
	CompanyAura: {
		Long:  CompanyNameLongAura,
		Short: CompanyNameShortAura,
	},
	CompanyBankCard: {
		Long:  CompanyNameLongBankcard,
		Short: CompanyNameShortBankcard,
	},
	CompanyCabal: {
		Long:  CompanyNameLongCabal,
		Short: CompanyNameShortCabal,
	},
	CompanyChinaUnionPay: {
		Long:  CompanyNameLongChinaUnionPay,
		Short: CompanyNameShortChinaUnionPay,
	},
	CompanyDankort: {
		Long:  CompanyNameLongDankort,
		Short: CompanyNameShortDankort,
	},
	CompanyDinersClubCarteBlance: {
		Long:  CompanyNameLongDinersClubCarteBlanche,
		Short: CompanyNameShortDinersClubCarteBlanche,
	},
	CompanyDinersClubEnRoute: {
		Long:  CompanyNameLongDinersClubEnRoute,
		Short: CompanyNameShortDinersClubEnRoute,
	},
	CompanyDinersClubInternational: {
		Long:  CompanyNameLongDinersClubInternational,
		Short: CompanyNameShortDinersClubInternational,
	},
	CompanyDiscover: {
		Long:  CompanyNameLongDiscover,
		Short: CompanyNameShortDiscover,
	},
	CompanyElo: {
		Long:  CompanyNameLongElo,
		Short: CompanyNameShortElo,
	},
	CompanyHipercard: {
		Long:  CompanyNameLongHipercard,
		Short: CompanyNameShortHipercard,
	},
	CompanyInstaPayment: {
		Long:  CompanyNameLongInstaPayment,
		Short: CompanyNameShortInstaPayment,
	},
	CompanyInterPayment: {
		Long:  CompanyNameLongInterPayment,
		Short: CompanyNameShortInterPayment,
	},
	CompanyJCB: {
		Long:  CompanyNameLongJCB,
		Short: CompanyNameShortJCB,
	},
	CompanyMaestro: {
		Long:  CompanyNameLongMaestro,
		Short: CompanyNameShortMaestro,
	},
	CompanyMasterCard: {
		Long:  CompanyNameLongMasterCard,
		Short: CompanyNameShortMasterCard,
	},
	CompanyNaranja: {
		Long:  CompanyNameLongNaranja,
		Short: CompanyNameShortNaranja,
	},
	CompanyVisa: {
		Long:  CompanyNameLongVisa,
		Short: CompanyNameShortVisa,
	},
	CompanyVisaElectron: {
		Long:  CompanyNameLongVisaElectron,
		Short: CompanyNameShortVisaElectron,
	},
}

func getCompany(name CompanyName) Company {
	return companies[name]
}
