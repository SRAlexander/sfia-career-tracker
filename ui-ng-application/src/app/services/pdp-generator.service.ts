import { Injectable } from '@angular/core';
import { RoleSkillsModel } from '../classes/role-skills-model';
import { SfiaSkillsService } from './sfia-skills.service';

@Injectable({
  providedIn: 'root'
})
export class PdpGeneratorService {

  constructor(private _skillService: SfiaSkillsService) { }

  public generatePDP(model : RoleSkillsModel) {

    const typeArray : any[] = [];

    let pdp = {
      roleTitle: model.jobTitle,
      sfiaLevel: model.sfiaLevel,
      coreSkills: typeArray,
      specialismSkills: typeArray
    }

    let coreSkills: any[] = []
    model.coreSkills.split(",").forEach(skillCode => {
      let foundSkills = this._skillService.getSFIASkillCriiteraByCodeAndLevel(skillCode, model.sfiaLevel);
      let skillTitle = this._skillService.getSFIASkills().filter((skill: { code: string; }) => skill.code === skillCode)[0];
      if (foundSkills == null || foundSkills.length === 0){
        foundSkills.push({
          description: "No SFIA criteria points have been found at this SFIA Level",
          skillCode: "NOTFOUND",
          keycode: 0,
          sfiaLevel: model.sfiaLevel
        })
      }

      if (typeof skillTitle !== 'undefined') {
        coreSkills = coreSkills.concat({skillsGroup: foundSkills, title: skillTitle.title})
      }
    })

    let specialismSkills: any[] = []
    model.specialismSkills.split(",").forEach(skillCode => {
      let foundSkills = this._skillService.getSFIASkillCriiteraByCodeAndLevel(skillCode, model.sfiaLevel);
      let skillTitle = this._skillService.getSFIASkills().filter((skill: { code: string; }) => skill.code === skillCode)[0];
      
      if (foundSkills == null || foundSkills.length === 0){
        foundSkills.push({
          description: "No SFIA criteria points have been found at this SFIA Level",
          skillCode: "NOTFOUND",
          keycode: 0,
          sfiaLevel: model.sfiaLevel
        })
      }
      if (typeof skillTitle !== 'undefined') {
        specialismSkills = specialismSkills.concat({skillsGroup: foundSkills, title: skillTitle.title})
      }
    })

    pdp.coreSkills = coreSkills;
    pdp.specialismSkills = specialismSkills;

    return pdp;
  }
}
