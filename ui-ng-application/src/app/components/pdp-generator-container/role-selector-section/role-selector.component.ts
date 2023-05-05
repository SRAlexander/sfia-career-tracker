import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { RolesService } from 'src/app/services/roles.service';
import { SfiaSkillsService } from 'src/app/services/sfia-skills.service';

@Component({
  selector: 'app-role-selector',
  templateUrl: './role-selector.component.html',
  styleUrls: ['./role-selector.component.scss']
})
export class RoleSelectorComponent implements OnInit {

  @Input() roleType: string = "";
  @Input() sfiaLevel: number = 0;
  @Output() onRoleSelected = new EventEmitter<any>();

  primaryRoles: any[] = [];

  constructor(
    private _rolesService: RolesService,
    private _sfiaSkillsService: SfiaSkillsService){
  }

  ngOnInit(): void {
    this.primaryRoles = [...new Set(this._rolesService.getRoleTypes())];
    
    const defaultRoleType = "Select a role type...";
    this.primaryRoles.splice(0, 0, defaultRoleType)
    if (this.roleType === "") {
      this.roleType = defaultRoleType;
    }

  }

  onRoleSelectedEvent(selectedRoleType: any) {
    let roleSearchResult = this._rolesService.getFullRoleBySFIAAndType(this.sfiaLevel, selectedRoleType)
    if (roleSearchResult.length > 0){
      roleSearchResult[0].coreSkills = this.replaceSkillCodeWithFullSkill(roleSearchResult[0].coreSkills);
      roleSearchResult[0].specialismSkills = this.replaceSkillCodeWithFullSkill(roleSearchResult[0].specialismSkills);
      this.onRoleSelected.emit(roleSearchResult[0]);
    } else {
      this.onRoleSelected.emit(null);
    }
  }

  replaceSkillCodeWithFullSkill(skillCodes: string[]) {
    const sfiaSkills = this._sfiaSkillsService.getSFIASkills();
    let fullSkills :any[] = [];
    skillCodes.forEach(skillCode => {
      fullSkills.push(sfiaSkills.filter((sfiaSkill : any) => sfiaSkill.code == skillCode)[0])
    })

    return fullSkills;
  }
}
