# SFIA Career Tracker

This project is designed to create a SFIA grading checklist based on an individual's target SFIA level and their required job skills. It is designed not only to manage SFIA-specific criteria but also to accommodate organisation-specific criteria. 

# Requirments

- A version of Go which can be downloaded from https://go.dev/dl/
- A copy of the SFIA Framework which can be acquired at https://sfia-online.org/en/sfia-8/documentation, an account is required.


# Orgnaisational Setup - Creating your criteria

As of creating this tool, the author does not feel that the SFIA-provided spreadsheet https://sfia-online.org/en/sfia-8/documentation is not in a suitable format for dynamic processing. Therefore there is a pre-step to run to extract all skills and responsibilities into a simple list.

Firstly, as an organisation, you will need to sign up for an account to download the SFIA framework spreadsheet, there is a license required at an organisational level, however, you can use a personal account to acquire the spreadsheet if you are testing this tool. You can then download the "sfia-8_en_200221.xlsx" as of 18/12/2022. Please place the file in the SFIA-RESOURCE folder. 

If you are using a newer version of the SFIA framework, this tool is designed to support future releases. You will however need to create some mapping files and modify some configurations. If you are using the version stated above, the applications are already configured. In the SFIA-RESOURCES folder, you can find "sfia-mapping-skills.txt" and "sfia-mapping-skills.txt". It contains a list of KEYS and COLUMNS required by the tool. The fields are as follows...

SHEETNAME: The name of the excel sheet to read from (not file)
SKILLCODE: The column letter associated with the skill code e.g. PROG. 
LEVEL-INDICATORS: A list of columns which indicate if a description exists for a given SFIA level starting at SFIA 1 and moving up through 2,3 etc... (NOTE: Responsibilities follow a different pattern so correspond to their rows and not columns)
LEVEL-DESCRIPTIONS: A list of columns which indicate where a description sits for a given SFIA level starting at SFIA 1 and moving up through 2,3 etc...

Finally, in the "app-critiera-extractor" folder, in the config file, set the name of the SFIA specification file against the variable "sfiaSpreadsheetLocation". You may also choose to output the processed data in a different format (JSON) but I would recommend sticking to EXCEL.

If you have the requirements above and have followed the steps you can run the following to create all SFIA criteria for all levels as a single Excel sheet...

```
cd app-criteria-extractor (go into the folder)
go run sfia-to-format.go
```

You should now find a timestamped file in the "sfia-processed-resources" folder.

If you have organisational-specific requirements against each SFIA grade you can now add them to this excel sheet with your custom Skillcode e.g. your organisation's name. These additional requirements can be auto-generated into a role requirement in the second app.

# General Use - Creating role-specific criteria

To generate a list of role requirements you will need a processed SFIA file in your "sfia-processed-resources" folder. If you don't have one and your org is in control of your SFIA requirements, ask someone to create a file. If you do have the file, head into the "app-pdp-criteria-generator" folder and open the config.json file. Make sure the "sfiaProcessedSpreadsheetLocation" parameter matches the name of your file. 

If you have any custom org requirements I would also recommend adding them to the "defaultSkills" parameter by adding the Skillcode you assigned them to the list separated with a comma.

To create the SFIA-level requirements run the following command.

```
cd app-pdp-criteria-generator (go into the folder)
go run create-pdp-criteria.go -sfia-level x CORE
```

Where x is the SFIA level you are generating against. 
This command will pull out all SFIA requirements for an individual acting at SFIA level x for all "defaultSkills" defined in the config.json. The generated file can be found in the root directory of the project.

To generate additional skill requirements for a specific role for example for a software developer in a data engineering role. You can run the following...

```
go run create-pdp-criteria.go -sfia-level 4 CORE PROG DENG
```

This will create a full job spec of requirements against SFIA. 

The tool also allows skill-only lookups by removing "CORE" from the command e.g.

```
go run create-pdp-criteria.go -sfia-level 3 PROG
```

A full list of supported skills as of SFIA Version 8 can be found below. 

# Skills

Acceptance testing                          BPTS
Animation development                       ADEV
Application support                         ASUP
Asset management                            ASMG
Audit                                       AUDT
Availability management                     AVMT
Benefits management                         BENM
Business administration                     ADMN
Business intelligence                       BINT
Business modelling                          BSMO
Business process improvement                BPRE
Business situation analysis                 BUSA
Capacity management                         CPMG
Certification scheme operation              CSOP
Change control                              CHMG
Competency assessment                       LEDA
Configuration management                    CFMG
Consultancy                                 CNSL
Content authoring                           INCA
Content publishing                          ICPM
Continuity management                       COPL
Contract management                         ITCM
Customer service support                    CSMG
Data engineering                            DENG
Data management                             DATM
Data modelling and design                   DTAN
Data science                                DATS
Data visualisation                          VISL
Database administration                     DBAD
Database design                             DBDS
Demand management                           DEMM
Digital forensics                           DGFS
Emerging technology monitoring              EMRG
Employee experience                         EEXP
Enterprise and business architecture        STPL
Facilities management                       DCMA
Feasibility assessment                      FEAS
Financial management                        FMIT
Governance                                  GOVN
Hardware design                             HWDE
High-performance computing                  HPCC
Incident management                         USUP
Information assurance                       INAS
Information management                      IRMG
Information security                        SCTY
Information systems coordination            ISCO
Innovation                                  INOV
Investment appraisal                        INVA
IT infrastructure                           ITOP
Knowledge management                        KNOW
Learning and development management         ETMG
Learning delivery                           ETDL
Learning design and development             TMCR
Machine learning                            MLNG
Marketing                                   MKTG
Measurement                                 MEAS
Methods and tools                           METL
Network design                              NTDS
Network support                             NTAS
Numerical analysis                          NUAN
Organisation design and implementation      ORDI
Organisational capability development       OCDV
Organisational change management            CIPM
Organisational facilitation                 OFCL
Penetration testing                         PENT
Performance management                      PEMT
Personal data protection                    PEDP
Portfolio management                        POMG
Portfolio, programme and project support    PROF
Problem management                          PBMG
Product management                          PROD
Professional development                    PDSV
Programme management                        PGMG
Programming/software development            PROG
Project management                          PRMG
Quality assurance                           QUAS
Quality management                          QUMG
Radio frequency engineering                 RFEN
Real-time/embedded systems development      RESD
Release and deployment                      RELM
Requirements definition and management      REQM
Research                                    RSCH
Resourcing                                  RESC
Risk management                             BURM
Safety assessment                           SFAS
Safety engineering                          SFEN
Sales support                               SSUP
Scientific modelling                        SCMO
Security operations                         SCAD
Selling                                     SALE
Service acceptance                          SEAC
Service catalogue management                SCMG
Service level management                    SLMO
Software configuration                      PORT
Software design                             SWDN
Solution architecture                       ARCH
Sourcing                                    SORC
Specialist advice                           TECH
Stakeholder relationship management         RLMT
Storage management                          STMG
Strategic planning                          ITSP
Subject formation                           SUBF
Supplier management                         SUPP
Sustainability                              SUST
System software                             SYSP
Systems and software life cycle engineering SLEN
Systems design                              DESN
Systems development management              DLMG
Systems installation and removal            HSIN
Systems integration and build               SINT
Teaching                                    TEAC
Technology service management               ITMG
Testing                                     TEST
Threat intelligence                         THIN
User experience analysis                    UNAN
User experience design                      HCEV
User experience evaluation                  USEV
User research                               URCH
Vulnerability assessment                    VUAS
Vulnerability research                      VURE
Workforce planning                          WFPL

