package main

import (
	"encoding/json"
	"fmt"
)

type VAConfig struct {
	Configuration struct {
		VATitle struct {
			Title       string `json:"Title,omitempty"`
			Description string `json:"Description,omitempty"`
			HoverOver   string `json:"HoverOver,omitempty"`
			SubTitle    string `json:"SubTitle,omitempty"`
		} `json:"VA_Title,omitempty"`
		VAAnnualGrowthRate struct {
			Label struct {
				Text string `json:"Text,omitempty"`
			} `json:"Label,omitempty"`
			ID       string `json:"ID,omitempty"`
			DataType string `json:"DataType,omitempty"`
			Field    struct {
				Type        string    `json:"Type,omitempty"`
				Values      []float64 `json:"Values,omitempty"`
				HoverOverEx struct {
					HoverOver string `json:"HoverOver,omitempty"`
					Replace   string `json:"Replace,omitempty"`
				} `json:"HoverOverEx,omitempty"`
			} `json:"Field,omitempty"`
		} `json:"VA_Annual_Growth_Rate,omitempty"`
		VARevenue         map[string]interface{} `json:"VA_Revenue,omitempty"`
		VAProductResale   map[string]interface{} `json:"VA_ProductResale,omitempty"`
		VATMSupport       map[string]interface{} `json:"VA_TM_Support,omitempty"`
		VAProjects        map[string]interface{} `json:"VA_Projects,omitempty"`
		VAManagedServices map[string]interface{} `json:"VA_Managed_Services,omitempty"`
		VACloud           map[string]interface{} `json:"VA_Cloud,omitempty"`
		VATotal           map[string]interface{} `json:"VA_Total,omitempty"`
		VANetIncome       map[string]interface{} `json:"VA_NetIncome,omitempty"`
		Inputs            map[string]interface{} `json:"Inputs,omitempty"`
	} `json:"Assessment,omitempty"`

	Dependencies struct {
		VaRevenueProductResale   []string `json:"VA_REVENUE_PRODUCT_RESALE,omitempty"`
		VaRevenueTmSupport       []string `json:"VA_REVENUE_TM_SUPPORT,omitempty"`
		VaRevenueProjects        []string `json:"VA_REVENUE_PROJECTS,omitempty"`
		VaRevenueManagedServices []string `json:"VA_REVENUE_MANAGED_SERVICES,omitempty"`
		VaRevenueCloud           []string `json:"VA_REVENUE_CLOUD,omitempty"`
	} `json:"Dependencies,omitempty"`

	Constants struct {
		FiveYEAR                   string  `json:"5YEAR,omitempty"`
		VaSlinsightsPbm            string  `json:"VA_SLINSIGHTS_PBM,omitempty"`
		VaSlinsightsPbmType        string  `json:"VA_SLINSIGHTS_PBM_TYPE,omitempty"`
		VaProductResaleMultplier   float64 `json:"VA_PRODUCT_RESALE_MULTPLIER,omitempty"`
		VaTmSupportMultplier       float64 `json:"VA_TM_SUPPORT_MULTPLIER,omitempty"`
		VaProjectsMultplier        float64 `json:"VA_PROJECTS_MULTPLIER,omitempty"`
		VaManagedServicesMultplier float64 `json:"VA_MANAGED_SERVICES_MULTPLIER,omitempty"`
		VaCloudMultplier           float64 `json:"VA_CLOUD_MULTPLIER,omitempty"`
	} `json:"Constants,omitempty"`

	Calculations struct {
		Section1BusinessPerformanceOverTheLast12Months struct {
			Step2ValuationsByRevenueAndNetIncome struct {
				VaRevenueTotal             string `json:"VA_REVENUE_TOTAL,omitempty"`
				VaProductResaleValuation   string `json:"VA_PRODUCT_RESALE_VALUATION,omitempty"`
				VaTmSupportValuation       string `json:"VA_TM_SUPPORT_VALUATION,omitempty"`
				VaProjectsValuation        string `json:"VA_PROJECTS_VALUATION,omitempty"`
				VaManagedServicesValuation string `json:"VA_MANAGED_SERVICES_VALUATION,omitempty"`
				VaCloudValuation           string `json:"VA_CLOUD_VALUATION,omitempty"`
				VaRevenueTotalValuation    string `json:"VA_REVENUE_TOTAL_VALUATION,omitempty"`
				VaNetIncome                string `json:"VA_NET_INCOME,omitempty"`
				VaNetIncomeMultiplier      struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
				} `json:"VA_NET_INCOME_MULTIPLIER,omitempty"`
				VaNetIncomeValuation string `json:"VA_NET_INCOME_VALUATION,omitempty"`
			} `json:"Step 2 : Valuations by Revenue and Net Income,omitempty"`
			Step3AverageOfRevenueAndNetIncomeValuations struct {
				VaAverageRevenueNetIncomeValuation string `json:"VA_AVERAGE_REVENUE_NET_INCOME_VALUATION,omitempty"`
			} `json:"Step 3 : Average of Revenue and Net Income valuations,omitempty"`
		} `json:"Section 1: Business Performance over the last 12 Months,omitempty"`
		Section2BusinessPerformanceIn5Yrs struct {
			Step5ProjectedRevenueInNext5Years struct {
				VaRevenueProductResalePc    string `json:"VA_REVENUE_PRODUCT_RESALE_PC,omitempty"`
				VaRevenueTmSupportPc        string `json:"VA_REVENUE_TM_SUPPORT_PC,omitempty"`
				VaRevenueProjectsPc         string `json:"VA_REVENUE_PROJECTS_PC,omitempty"`
				VaRevenueManagedServicesPc  string `json:"VA_REVENUE_MANAGED_SERVICES_PC,omitempty"`
				VaRevenueCloudPc            string `json:"VA_REVENUE_CLOUD_PC,omitempty"`
				VaRevenueTotal1Yr           string `json:"VA_REVENUE_TOTAL_1YR,omitempty"`
				VaRevenueTotal2Yr           string `json:"VA_REVENUE_TOTAL_2YR,omitempty"`
				VaRevenueTotal3Yr           string `json:"VA_REVENUE_TOTAL_3YR,omitempty"`
				VaRevenueTotal4Yr           string `json:"VA_REVENUE_TOTAL_4YR,omitempty"`
				VaRevenueTotal5Yr           string `json:"VA_REVENUE_TOTAL_5YR,omitempty"`
				VaRevenueProductResale5Yr   string `json:"VA_REVENUE_PRODUCT_RESALE_5YR,omitempty"`
				VaRevenueTmSupport5Yr       string `json:"VA_REVENUE_TM_SUPPORT_5YR,omitempty"`
				VaRevenueProjects5Yr        string `json:"VA_REVENUE_PROJECTS_5YR,omitempty"`
				VaRevenueManagedServices5Yr string `json:"VA_REVENUE_MANAGED_SERVICES_5YR,omitempty"`
				VaRevenueCloud5Yr           string `json:"VA_REVENUE_CLOUD_5YR,omitempty"`
				VaRevenueTotalMm            string `json:"VA_REVENUE_TOTAL_MM,omitempty"`
				VaRevenueTotal1YrMm         string `json:"VA_REVENUE_TOTAL_1YR_MM,omitempty"`
				VaRevenueTotal2YrMm         string `json:"VA_REVENUE_TOTAL_2YR_MM,omitempty"`
				VaRevenueTotal3YrMm         string `json:"VA_REVENUE_TOTAL_3YR_MM,omitempty"`
				VaRevenueTotal4YrMm         string `json:"VA_REVENUE_TOTAL_4YR_MM,omitempty"`
				VaRevenueTotal5YrMm         string `json:"VA_REVENUE_TOTAL_5YR_MM,omitempty"`
				VaProductResaleMultplier5Yr struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
					Value     string `json:"Value,omitempty"`
				} `json:"VA_PRODUCT_RESALE_MULTPLIER_5YR,omitempty"`
				VaTmSupportMultplier5Yr struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
					Value     string `json:"Value,omitempty"`
				} `json:"VA_TM_SUPPORT_MULTPLIER_5YR,omitempty"`
				VaProjectsMultplier5Yr struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
					Value     string `json:"Value,omitempty"`
				} `json:"VA_PROJECTS_MULTPLIER_5YR,omitempty"`
				VaManagedServicesMultplier5Yr struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
					Value     string `json:"Value,omitempty"`
				} `json:"VA_MANAGED_SERVICES_MULTPLIER_5YR,omitempty"`
				VaCloudMultplier5Yr struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
					Value     string `json:"Value,omitempty"`
				} `json:"VA_CLOUD_MULTPLIER_5YR,omitempty"`
				VaProductResaleValuation5Yr   string `json:"VA_PRODUCT_RESALE_VALUATION_5YR,omitempty"`
				VaTmSupportValuation5Yr       string `json:"VA_TM_SUPPORT_VALUATION_5YR,omitempty"`
				VaProjectsValuation5Yr        string `json:"VA_PROJECTS_VALUATION_5YR,omitempty"`
				VaManagedServicesValuation5Yr string `json:"VA_MANAGED_SERVICES_VALUATION_5YR,omitempty"`
				VaCloudValuation5Yr           string `json:"VA_CLOUD_VALUATION_5YR,omitempty"`
				VaRevenueTotalValuation5Yr    string `json:"VA_REVENUE_TOTAL_VALUATION_5YR,omitempty"`
				VaRevenueTotal5YrCagr         struct {
					Operation    string `json:"Operation,omitempty"`
					Nper         int    `json:"NPER,omitempty"`
					PresentValue string `json:"PresentValue,omitempty"`
					FutureValue  string `json:"FutureValue,omitempty"`
				} `json:"VA_REVENUE_TOTAL_5YR_CAGR,omitempty"`
				VaNetIncome5Yr         string `json:"VA_NET_INCOME_5YR,omitempty"`
				VaNetIncomeMultiple5Yr struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
				} `json:"VA_NET_INCOME_MULTIPLE_5YR,omitempty"`
				VaNetIncomeValuation5Yr     string `json:"VA_NET_INCOME_VALUATION_5YR,omitempty"`
				VaNetIncomeValuation5YrCagr struct {
					Operation    string `json:"Operation,omitempty"`
					Nper         int    `json:"NPER,omitempty"`
					PresentValue string `json:"PresentValue,omitempty"`
					FutureValue  string `json:"FutureValue,omitempty"`
				} `json:"VA_NET_INCOME_VALUATION_5YR_CAGR,omitempty"`
				VaAccountsCagr struct {
					Operation    string `json:"Operation,omitempty"`
					Nper         int    `json:"NPER,omitempty"`
					PresentValue string `json:"PresentValue,omitempty"`
					FutureValue  string `json:"FutureValue,omitempty"`
				} `json:"VA_ACCOUNTS_CAGR,omitempty"`
			} `json:"Step 5 : Projected Revenue in  next 5 years,omitempty"`
			Step6AverageOfFutureRevenueAndNetIncomeValuations struct {
				VaAverageRevenueNetIncomeValuation5Yr     string `json:"VA_AVERAGE_REVENUE_NET_INCOME_VALUATION_5YR,omitempty"`
				VaAverageRevenueNetIncomeValuation5YrCagr struct {
					Operation    string `json:"Operation,omitempty"`
					Nper         int    `json:"NPER,omitempty"`
					PresentValue string `json:"PresentValue,omitempty"`
					FutureValue  string `json:"FutureValue,omitempty"`
				} `json:"VA_AVERAGE_REVENUE_NET_INCOME_VALUATION_5YR_CAGR,omitempty"`
			} `json:"Step6 : Average of future Revenue and Net Income valuations,omitempty"`
		} `json:"Section 2: Business Performance in 5 yrs,omitempty"`
		Section3SalesRampUpNeededToAttainBusinessPerformanceAsPerPlan struct {
			Step7AverageCusomterSizeAndAnnualAccountAttrition struct {
				VaRevenuePerCustomer  string `json:"VA_REVENUE_PER_CUSTOMER,omitempty"`
				VaAvgCustomersCurrent string `json:"VA_AVG_CUSTOMERS_CURRENT,omitempty"`
			} `json:"Step7 : Average cusomter size and annual account attrition,omitempty"`
			Step8SalesRampUpDecidePracticalityAdjustAsNeeded struct { //nolint
				VaAccounts    string `json:"VA_ACCOUNTS,omitempty"`
				VaAccounts1Yr string `json:"VA_ACCOUNTS_1YR,omitempty"`
				VaAccounts2Yr string `json:"VA_ACCOUNTS_2YR,omitempty"`
				VaAccounts3Yr string `json:"VA_ACCOUNTS_3YR,omitempty"`
				VaAccounts4Yr string `json:"VA_ACCOUNTS_4YR,omitempty"`
				VaAccounts5Yr string `json:"VA_ACCOUNTS_5YR,omitempty"`
			} `json:"Step8: Sales ramp up,decide practicality & adjust as needed,omitempty"` //nolint
		} `json:"Section 3: Sales ramp up needed to attain business performance as per plan,omitempty"`
	} `json:"Calculations,omitempty"`
	LookupTables struct {
		VaUsersMedianTable struct {
			Five25Users   int `json:"5-25 Users,omitempty"`
			Two6100Users  int `json:"26-100 Users,omitempty"`
			One01500Users int `json:"101-500 Users,omitempty"`
			Five_00Users  int `json:">500 Users,omitempty"`
		} `json:"VA_USERS_MEDIAN_TABLE,omitempty"`
		VaRevenueValuationMultiple struct {
			Zero00  float64 `json:"0.00,omitempty"`
			Seven51 float64 `json:"7.51,omitempty"`
			One510  float64 `json:"15.10,omitempty"`
		} `json:"VA_REVENUE_VALUATION_MULTIPLE,omitempty"`
		VaRevenueMultiple5YrTable struct {
			A30LowerThanToday struct {
				ProductResale   float64 `json:"Product Resale,omitempty"`
				TMSupport       float64 `json:"T&M Support,omitempty"`
				Projects        float64 `json:"Projects,omitempty"`
				ManagedServices float64 `json:"Managed Services,omitempty"`
				YourOwnCloud    float64 `json:"Your Own Cloud,omitempty"`
			} `json:"a) 30% lower than today,omitempty"`
			BAboutTheSameAsToday struct {
				ProductResale   float64 `json:"Product Resale,omitempty"`
				TMSupport       float64 `json:"T&M Support,omitempty"`
				Projects        float64 `json:"Projects,omitempty"`
				ManagedServices float64 `json:"Managed Services,omitempty"`
				YourOwnCloud    float64 `json:"Your Own Cloud,omitempty"`
			} `json:"b) About the same as today,omitempty"`
			C30HigherThanToday struct {
				ProductResale   float64 `json:"Product Resale,omitempty"`
				TMSupport       float64 `json:"T&M Support,omitempty"`
				Projects        float64 `json:"Projects,omitempty"`
				ManagedServices float64 `json:"Managed Services,omitempty"`
				YourOwnCloud    float64 `json:"Your Own Cloud,omitempty"`
			} `json:"c) 30% higher than today,omitempty"`
		} `json:"VA_REVENUE_MULTIPLE_5YR_TABLE,omitempty"`
		VaRevenueEmployeesServiceCentricTable struct {
			Num5 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"-5,omitempty"`
			Five10 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"5-10,omitempty"`
			One020 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"10-20,omitempty"`
			Two0 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"20-,omitempty"`
		} `json:"VA_REVENUE_EMPLOYEES_SERVICE_CENTRIC_TABLE,omitempty"`
		VaRevenueEmployeesProductCentricTable struct {
			Num1 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"-1,omitempty"`
			One5 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   int     `json:"% Sales Employees,omitempty"`
			} `json:"1-5,omitempty"`
			Five10 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"5-10,omitempty"`
			One020 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"10-20,omitempty"`
			Two030 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"20-30,omitempty"`
			Three050 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   int     `json:"% Sales Employees,omitempty"`
			} `json:"30-50,omitempty"`
			Five0100 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"50-100,omitempty"`
			One00 struct {
				ServiceEmployees float64 `json:"% Service Employees,omitempty"`
				SalesEmployees   float64 `json:"% Sales Employees,omitempty"`
			} `json:"100-,omitempty"`
		} `json:"VA_REVENUE_EMPLOYEES_PRODUCT_CENTRIC_TABLE,omitempty"`
	} `json:"LookupTables,omitempty"`
	Reports struct {
		Heading struct {
			Title       string `json:"Title,omitempty"`
			Description string `json:"Description,omitempty"`
			SubTitle    struct {
				Text    string `json:"Text,omitempty"`
				Replace string `json:"Replace,omitempty"`
			} `json:"Sub-Title,omitempty"`
		} `json:"Heading,omitempty"`
		RevenueProfitValuation struct {
			VizType string `json:"Viz Type,omitempty"`
			Title   string `json:"Title,omitempty"`
			Columns []struct {
				Label struct {
					Text string `json:"Text,omitempty"`
				} `json:"Label,omitempty"`
				HoverOver string `json:"Hover Over,omitempty"`
				LabelEx   struct {
					Label   string `json:"Label,omitempty"`
					Replace string `json:"Replace,omitempty"`
				} `json:"LabelEx,omitempty"`
			} `json:"Columns,omitempty"`
			Rows     [][]any `json:"Rows,omitempty"`
			Footnote string  `json:"Footnote,omitempty"`
		} `json:"Revenue_Profit_Valuation,omitempty"`
		Employees struct {
			VizType    string   `json:"Viz Type,omitempty"`
			Title      string   `json:"Title,omitempty"`
			XAxisTicks []string `json:"X-Axis Ticks,omitempty"`
			YAxisTicks []int    `json:"Y-Axis Ticks,omitempty"`
			YAxisData  struct {
				ServiceHeadcount string `json:"Service Headcount,omitempty"`
				SalesHeadcount   string `json:"Sales Headcount,omitempty"`
				OtherHeadcount   string `json:"Other Headcount,omitempty"`
			} `json:"Y-Axis Data,omitempty"`
			YAxisLabel   string `json:"Y-Axis Label,omitempty"`
			Legend       bool   `json:"Legend,omitempty"`
			Calculations struct {
				VaTotalEmployees   string `json:"VA_TOTAL_EMPLOYEES,omitempty"`
				VASERVICEEMPLOYEES struct {
					Type      string `json:"Type,omitempty"`
					Operation string `json:"Operation,omitempty"`
					TableEx   struct {
						Table   string `json:"Table,omitempty"`
						Replace string `json:"Replace,omitempty"`
					} `json:"TableEx,omitempty"`
					Keys   []string `json:"Keys,omitempty"`
					Values string   `json:"Values,omitempty"`
				} `json:"VA_%_SERVICE_EMPLOYEES,omitempty"`
				VASALESEMPLOYEES struct {
					Type      string `json:"Type,omitempty"`
					Operation string `json:"Operation,omitempty"`
					TableEx   struct {
						Table   string `json:"Table,omitempty"`
						Replace string `json:"Replace,omitempty"`
					} `json:"TableEx,omitempty"`
					Keys   []string `json:"Keys,omitempty"`
					Values string   `json:"Values,omitempty"`
				} `json:"VA_%_SALES_EMPLOYEES,omitempty"`
				VaServiceHeadcount struct {
					Type    string `json:"Type,omitempty"`
					Type2   string `json:"Type-2,omitempty"`
					Formula string `json:"Formula,omitempty"`
				} `json:"VA_SERVICE_HEADCOUNT,omitempty"`
				VaSalesHeadcount struct {
					Type    string `json:"Type,omitempty"`
					Type2   string `json:"Type-2,omitempty"`
					Formula string `json:"Formula,omitempty"`
				} `json:"VA_SALES_HEADCOUNT,omitempty"`
				VaOtherHeadcount struct {
					Type    string `json:"Type,omitempty"`
					Type2   string `json:"Type-2,omitempty"`
					Formula string `json:"Formula,omitempty"`
				} `json:"VA_OTHER_HEADCOUNT,omitempty"`
			} `json:"Calculations,omitempty"`
		} `json:"Employees,omitempty"`
		NewAccounts struct {
			VizType []string `json:"Viz Type,omitempty"`
			Title   string   `json:"Title,omitempty"`
			XAxis   []string `json:"X-Axis,omitempty"`
			YAxis   struct {
				YearEndOfAccounts []string `json:"Year End # of Accounts,omitempty"`
				AccountsAttrited  []string `json:"Accounts Attrited,omitempty"`
				WinsNeededPerYear []string `json:"# Wins Needed per Year,omitempty"`
				YAxisLabel        string   `json:"Y-Axis Label,omitempty"`
			} `json:"Y-Axis,omitempty"`
			Calculations struct {
				VaAccountsAttrited1Yr string `json:"VA_ACCOUNTS_ATTRITED_1YR,omitempty"`
				VaAccountsAttrited2Yr string `json:"VA_ACCOUNTS_ATTRITED_2YR,omitempty"`
				VaAccountsAttrited3Yr string `json:"VA_ACCOUNTS_ATTRITED_3YR,omitempty"`
				VaAccountsAttrited4Yr string `json:"VA_ACCOUNTS_ATTRITED_4YR,omitempty"`
				VaAccountsAttrited5Yr string `json:"VA_ACCOUNTS_ATTRITED_5YR,omitempty"`
				VaWinsNeeded1Yr       string `json:"VA_WINS_NEEDED_1YR,omitempty"`
				VaWinsNeeded2Yr       string `json:"VA_WINS_NEEDED_2YR,omitempty"`
				VaWinsNeeded3Yr       string `json:"VA_WINS_NEEDED_3YR,omitempty"`
				VaWinsNeeded4Yr       string `json:"VA_WINS_NEEDED_4YR,omitempty"`
				VaWinsNeeded5Yr       string `json:"VA_WINS_NEEDED_5YR,omitempty"`
			} `json:"Calculations,omitempty"`
		} `json:"New_Accounts,omitempty"`
		UsersSupported struct {
			VizType      string   `json:"Viz Type,omitempty"`
			Title        string   `json:"Title,omitempty"`
			XAxisTicks   []string `json:"X-Axis Ticks,omitempty"`
			YAxisTicks   []any    `json:"Y-Axis Ticks,omitempty"`
			YAxisData    []string `json:"Y-Axis Data,omitempty"`
			YAxisLabel   string   `json:"Y-Axis Label,omitempty"`
			Calculations struct {
				VaRevenueGap     string `json:"VA_REVENUE_GAP,omitempty"`
				VaAccountsNeeded string `json:"VA_ACCOUNTS_NEEDED,omitempty"`
				VaUsersMedian    struct {
					Operation string `json:"Operation,omitempty"`
					Table     string `json:"Table,omitempty"`
					Key       string `json:"Key,omitempty"`
				} `json:"VA_USERS_MEDIAN,omitempty"`
				VaTotalUsers    string `json:"VA_TOTAL_USERS,omitempty"`
				VaTotalUsers1Yr string `json:"VA_TOTAL_USERS_1YR,omitempty"`
				VaTotalUsers2Yr string `json:"VA_TOTAL_USERS_2YR,omitempty"`
				VaTotalUsers3Yr string `json:"VA_TOTAL_USERS_3YR,omitempty"`
				VaTotalUsers4Yr string `json:"VA_TOTAL_USERS_4YR,omitempty"`
				VaTotalUsers5Yr string `json:"VA_TOTAL_USERS_5YR,omitempty"`
			} `json:"Calculations,omitempty"`
		} `json:"Users_Supported,omitempty"`
		SLInsights struct {
			VizType string `json:"Viz Type,omitempty"`
			Title   string `json:"Title,omitempty"`
			TextEx  struct {
				Text    []string `json:"Text,omitempty"`
				Replace []string `json:"Replace,omitempty"`
			} `json:"TextEx,omitempty"`
			Calculations struct {
				VaWinsNeededAverage string `json:"VA_WINS_NEEDED_AVERAGE,omitempty"`
			} `json:"Calculations,omitempty"`
		} `json:"S-L_Insights,omitempty"`
	} `json:"Reports,omitempty"`
}

func main() {
	if d, e := json.Marshal(VAConfig{}); e != nil {
		fmt.Print(e)
	} else {
		fmt.Print(string(d))
	}
}
