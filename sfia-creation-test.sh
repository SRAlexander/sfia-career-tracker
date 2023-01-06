cd app-pdp-criteria-generator
go run create-pdp-criteria.go --sfia-level 1 --output-filename softeng-sfia-1.xlsx CORE PROG
go run create-pdp-criteria.go --sfia-level 2 --output-filename softeng-sfia-2.xlsx CORE PROG
go run create-pdp-criteria.go --sfia-level 3 --output-filename softeng-sfia-3.xlsx CORE PROG
go run create-pdp-criteria.go --sfia-level 4 --output-filename softeng-sfia-4.xlsx CORE PROG
go run create-pdp-criteria.go --sfia-level 5 --output-filename softeng-sfia-5.xlsx CORE PROG
go run create-pdp-criteria.go --sfia-level 6 --output-filename softeng-sfia-6.xlsx CORE PROG
go run create-pdp-criteria.go --sfia-level 7 --output-filename softeng-sfia-7.xlsx CORE PROG

cd ..
cd app-pdp-generator

go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-1.xlsx --output-filename softeng-sfia-1-pdp.MD
go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-2.xlsx --output-filename softeng-sfia-2-pdp.MD
go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-3.xlsx --output-filename softeng-sfia-3-pdp.MD
go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-4.xlsx --output-filename softeng-sfia-4-pdp.MD
go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-5.xlsx --output-filename softeng-sfia-5-pdp.MD 
go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-6.xlsx --output-filename softeng-sfia-6-pdp.MD 
go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-7.xlsx --output-filename softeng-sfia-7-pdp.MD 