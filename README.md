# SFIA Career Tracker

This project is designed to create a SFIA grading checklist based on an individuals target SFIA level and their required job skills. It is designed not only to manage SFIA specific critieria but also to accomodate orginsation specific critieria. 

# Orgnaisational Setup

As of creating this tool, the author does not feel that the SFIA provided spreadsheet https://sfia-online.org/en/sfia-8/documentation is not in a sutiable format for dynamic processing. Therefore there is a pre step to run to extract all skills and responsibilities into a simple list.

Firstly, as an orginsation, you will need to sign up for an account to download the sfia framework spreadsheet, there is a license required at an organisational level, however, you can use a persoanl account to aquire the spreadsheet if you are testing this tool. You can then download the "sfia-8_en_200221.xlsx" as of the 18/12/2022. Please place the file in the SFIA-RESOURCE folder. 

If you are using a newer version of the SFIA framework, this tool is designed to support future releases. You will however need to create some mapping files. In the SFIA-RESOURCES folder you can find a sfia-mapping-skills.txt and sfia-mapping-skills.txt. It contains a list of KEYS and COLUMNS required by the tool. The KEYS should be self explanatory as to what data they point to, update the respective column letter to where the data sits in the newer version.

go install github.com/xuri/excelize/v2@latest

# Skills

Acceptance testing	                        BPTS
Animation development	                    ADEV
Application support	                        ASUP
Asset management	                        ASMG
Audit	                                    AUDT
Availability management	                    AVMT
Benefits management	                        BENM
Business administration	                    ADMN
Business intelligence	                    BINT
Business modelling	                        BSMO
Business process improvement	            BPRE
Business situation analysis	                BUSA
Capacity management	                        CPMG
Certification scheme operation	            CSOP
Change control	                            CHMG
Competency assessment	                    LEDA
Configuration management	                CFMG
Consultancy	                                CNSL
Content authoring	                        INCA
Content publishing	                        ICPM
Continuity management	                    COPL
Contract management	                        ITCM
Customer service support	                CSMG
Data engineering	                        DENG
Data management	                            DATM
Data modelling and design	                DTAN
Data science	                            DATS
Data visualisation	                        VISL
Database administration	                    DBAD
Database design	                            DBDS
Demand management	                        DEMM
Digital forensics	                        DGFS
Emerging technology monitoring	            EMRG
Employee experience	                        EEXP
Enterprise and business architecture	    STPL
Facilities management	                    DCMA
Feasibility assessment	                    FEAS
Financial management	                    FMIT
Governance	                                GOVN
Hardware design	                            HWDE
High-performance computing	                HPCC
Incident management	                        USUP
Information assurance	                    INAS
Information management	                    IRMG
Information security	                    SCTY
Information systems coordination	        ISCO
Innovation	                                INOV
Investment appraisal	                    INVA
IT infrastructure	                        ITOP
Knowledge management	                    KNOW
Learning and development management	        ETMG
Learning delivery	                        ETDL
Learning design and development	            TMCR
Machine learning	                        MLNG
Marketing	                                MKTG
Measurement	                                MEAS
Methods and tools	                        METL
Network design	                            NTDS
Network support	                            NTAS
Numerical analysis	                        NUAN
Organisation design and implementation	    ORDI
Organisational capability development	    OCDV
Organisational change management	        CIPM
Organisational facilitation	                OFCL
Penetration testing	                        PENT
Performance management	                    PEMT
Personal data protection	                PEDP
Portfolio management	                    POMG
Portfolio, programme and project support	PROF
Problem management	                        PBMG
Product management	                        PROD
Professional development	                PDSV
Programme management	                    PGMG
Programming/software development	        PROG
Project management	                        PRMG
Quality assurance	                        QUAS
Quality management	                        QUMG
Radio frequency engineering	                RFEN
Real-time/embedded systems development	    RESD
Release and deployment	                    RELM
Requirements definition and management	    REQM
Research	                                RSCH
Resourcing	                                RESC
Risk management	                            BURM
Safety assessment	                        SFAS
Safety engineering	                        SFEN
Sales support	                            SSUP
Scientific modelling	                    SCMO
Security operations	                        SCAD
Selling	                                    SALE
Service acceptance	                        SEAC
Service catalogue management	            SCMG
Service level management	                SLMO
Software configuration	                    PORT
Software design	                            SWDN
Solution architecture	                    ARCH
Sourcing	                                SORC
Specialist advice	                        TECH
Stakeholder relationship management	        RLMT
Storage management	                        STMG
Strategic planning	                        ITSP
Subject formation	                        SUBF
Supplier management	                        SUPP
Sustainability	                            SUST
System software	                            SYSP
Systems and software life cycle engineering	SLEN
Systems design	                            DESN
Systems development management	            DLMG
Systems installation and removal	        HSIN
Systems integration and build	            SINT
Teaching	                                TEAC
Technology service management	            ITMG
Testing	                                    TEST
Threat intelligence	                        THIN
User experience analysis	                UNAN
User experience design	                    HCEV
User experience evaluation	                USEV
User research	                            URCH
Vulnerability assessment	                VUAS
Vulnerability research	                    VURE
Workforce planning	                        WFPL









