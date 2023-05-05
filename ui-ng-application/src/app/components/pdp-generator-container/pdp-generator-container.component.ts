import { Component } from '@angular/core';
import { PdpGeneratorService } from 'src/app/services/pdp-generator.service';

@Component({
  selector: 'app-pdp-generator-container',
  templateUrl: './pdp-generator-container.component.html',
  styleUrls: ['./pdp-generator-container.component.scss']
})

export class PdpGeneratorContainerComponent {

  constructor(
    private _pdpGeneratorService: PdpGeneratorService
  ){
  }

  sfiaLevel: number = 0; 
  creationType: string = "";
  jobTitle: string = "Custom Role"
  roleType: string = "";

  coreSkills: any[] = [{ "code":"AUTONOMY", "title":"Autonomy"},
                            { "code":"INFLUENCE", "title":"Influence"},
                            { "code":"COMPLEXITY", "title":"Complexity"},
                            { "code":"BUSINESSSKILLS", "title":"Business skills"},
                            { "code":"KNOWLEDGE", "title":"Knowledge"}];

  specialismSkills: any[] = [];
  generatedPdp: any = null;
  showPdp: boolean = false;

  onSFIASelected(selectedSfia : any) {
    this.sfiaLevel = selectedSfia;
  }

  onCreationTypeSelected(creationType: string) {
    this.creationType = creationType;

    if (this.creationType !== "roles") {
      this.jobTitle = "Custom Role"
    }
  }

  onRoleSelected(role : any){
    this.jobTitle = role.jobTitle;
    this.coreSkills = role.coreSkills;
    this.specialismSkills = role.specialismSkills;
    this.roleType = role.type;
  }

  generatePDP(){

    this.generatedPdp = this._pdpGeneratorService.generatePDP({
      jobTitle: this.jobTitle,
      coreSkills: this.coreSkills.map(skill => skill.code ?? skill).join(','),
      specialismSkills: this.specialismSkills.map(skill => skill.code ?? skill).join(','),
      sfiaLevel: this.sfiaLevel
    });

    this.showPdp = true;
  }

  returnFromViewer() {
    this.showPdp = false;
  }
}
