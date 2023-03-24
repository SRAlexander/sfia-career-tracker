import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { map } from 'rxjs/internal/operators/map';
import { RoleSkillsModel } from 'src/app/classes/role-skills-model';
import { SfiaSkillsService } from 'src/app/services/sfia-skills.service';

@Component({
  selector: 'app-skill-selector',
  templateUrl: './skill-selector.component.html',
  styleUrls: ['./skill-selector.component.scss']
})
export class SkillSelectorComponent implements OnInit {

  avaliableSkills: any[] = [];
  selectedCoreSkills: any[] = [{ "code":"AUTONOMY", "title":"Autonomy"},
                            { "code":"INFLUENCE", "title":"Influence"},
                            { "code":"COMPLEXITY", "title":"Complexity"},
                            { "code":"BUSINESSSKILLS", "title":"Business skills"},
                            { "code":"KNOWLEDGE", "title":"Knowledge"}];

  selectedSpecialismSkills: any[] = [];
  createdRole: RoleSkillsModel = new RoleSkillsModel();
  
  @Input() sfiaLevel: number = 0;
  @Output() onRoleCreated: EventEmitter<RoleSkillsModel> = new EventEmitter();

  constructor(private _sfiaSkillsService : SfiaSkillsService) {
  }

  ngOnInit(): void {

    this.avaliableSkills = this._sfiaSkillsService.getSFIASkills();

    let uniqueSkills: any[] = [];
    for (let avaliableSkillIndex = 0; avaliableSkillIndex < this.avaliableSkills.length; avaliableSkillIndex++){
      const selectedSkillSearch = this.selectedCoreSkills.filter(skill => skill.code === this.avaliableSkills[avaliableSkillIndex].code);
      if (selectedSkillSearch == null || selectedSkillSearch.length === 0) {
          uniqueSkills.push(this.avaliableSkills[avaliableSkillIndex])
        }
    }

    uniqueSkills.splice(0,0, {code: "NA", "title":"Select a skill to add it"});
    this.avaliableSkills = uniqueSkills;

    this.createdRole.sfiaLevel = this.sfiaLevel;
    this.createdRole.jobTitle = "Generated Role"
    this.updateCreatedRole();
  }

  removeCoreSkill(skillIndex: number) {
    this.avaliableSkills.push(this.selectedCoreSkills[skillIndex])
    this.selectedCoreSkills.splice(skillIndex,1)
    this.updateCreatedRole();
  }

  addCoreSkill(event: any) {
    var selectedSkillTitle = event.target.value;
    var selectedSkills = this.avaliableSkills.filter(skill => skill.title === selectedSkillTitle);

    if (selectedSkills.length >0 && this.selectedCoreSkills.indexOf(selectedSkills[0]) === -1 && selectedSkills[0].code !== "NA") {
      this.selectedCoreSkills.push(selectedSkills[0])
      this.avaliableSkills = this.avaliableSkills.filter(skill => skill !== selectedSkills[0])
    }

    this.updateCreatedRole();
  }

  removeSpecialismSkill(skillIndex: number) {
    this.avaliableSkills.push(this.selectedSpecialismSkills[skillIndex])
    this.selectedSpecialismSkills.splice(skillIndex,1)
    this.updateCreatedRole();
  }

  addSpecialismSkill(event: any) {
    var selectedSkillTitle = event.target.value;
    var selectedSkills = this.avaliableSkills.filter(skill => skill.title === selectedSkillTitle);

    if (selectedSkills.length >0 && this.selectedSpecialismSkills.indexOf(selectedSkills[0]) === -1 && selectedSkills[0].code !== "NA") {
      this.selectedSpecialismSkills.push(selectedSkills[0])
      this.avaliableSkills = this.avaliableSkills.filter(skill => skill !== selectedSkills[0])
    }

    this.updateCreatedRole();
  }

  updateCreatedRole() {
    this.createdRole.coreSkills = this.selectedCoreSkills.map(skill => skill.code).join(",");
    this.createdRole.specialismSkills = this.selectedSpecialismSkills.map(skill => skill.code).join(",")
    this.onRoleCreated.emit(this.createdRole);
  }
}
