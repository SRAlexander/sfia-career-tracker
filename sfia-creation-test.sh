cd app-pdp-criteria-generator
# go run create-pdp-criteria.go --sfia-level 1 --output-filename softeng-sfia-1.xlsx CORE PROG
# go run create-pdp-criteria.go --sfia-level 2 --output-filename softeng-sfia-2.xlsx CORE PROG BUSA
# go run create-pdp-criteria.go --sfia-level 3 --output-filename softeng-sfia-3.xlsx CORE PROG
# go run create-pdp-criteria.go --sfia-level 4 --output-filename softeng-sfia-4.xlsx CORE PROG
# go run create-pdp-criteria.go --sfia-level 5 --output-filename softeng-sfia-5.xlsx CORE PROG
# go run create-pdp-criteria.go --sfia-level 6 --output-filename softeng-sfia-6.xlsx CORE PROG
# go run create-pdp-criteria.go --sfia-level 7 --output-filename softeng-sfia-7.xlsx CORE PROG

# go run create-pdp-criteria.go --sfia-level 5 --output-filename softengdata-sfia-5.xlsx CORE PROG MLNG DENG
# go run create-pdp-criteria.go --sfia-level 1 --output-filename ba-sfia-1.xlsx CORE BUSA
# go run create-pdp-criteria.go --sfia-level 2 --output-filename ba-sfia-2.xlsx CORE BUSA
# go run create-pdp-criteria.go --sfia-level 3 --output-filename ba-sfia-3.xlsx CORE BUSA
# go run create-pdp-criteria.go --sfia-level 4 --output-filename ba-sfia-4.xlsx CORE BUSA
# go run create-pdp-criteria.go --sfia-level 5 --output-filename ba-sfia-5.xlsx CORE BUSA
# go run create-pdp-criteria.go --sfia-level 6 --output-filename ba-sfia-6.xlsx CORE BUSA
# go run create-pdp-criteria.go --sfia-level 7 --output-filename ba-sfia-7.xlsx CORE BUSA
# go run create-pdp-criteria.go --sfia-level 7 --output-filename ba-sfia-7.xlsx CORE BUSA

go run create-pdp-criteria.go --sfia-level 5 --output-filename pt-sfia-5.xlsx CORE ORDI OCDV MTEK

cd ..
cd app-pdp-generator

# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-1.xlsx --output-filename softeng-sfia-1-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-2.xlsx --output-filename softeng-sfia-2-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-3.xlsx --output-filename softeng-sfia-3-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-4.xlsx --output-filename softeng-sfia-4-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-5.xlsx --output-filename softeng-sfia-5-pdp.MD 
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-6.xlsx --output-filename softeng-sfia-6-pdp.MD 
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softeng-sfia-7.xlsx --output-filename softeng-sfia-7-pdp.MD 

go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/pt-sfia-5.xlsx --output-filename pt-sfia-5-pdp.MD 


go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/softengdata-sfia-5.xlsx --output-filename softengdata-sfia-5-pdp.MD 

# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/ba-sfia-1.xlsx --output-filename ba-sfia-1-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/ba-sfia-2.xlsx --output-filename ba-sfia-2-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/ba-sfia-3.xlsx --output-filename ba-sfia-3-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/ba-sfia-4.xlsx --output-filename ba-sfia-4-pdp.MD
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/ba-sfia-5.xlsx --output-filename ba-sfia-5-pdp.MD 
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/ba-sfia-6.xlsx --output-filename ba-sfia-6-pdp.MD 
# go run pdp-generator.go --skill-list ../app-pdp-criteria-generator/output/ba-sfia-7.xlsx --output-filename ba-sfia-7-pdp.MD 