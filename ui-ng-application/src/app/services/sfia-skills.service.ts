import { Injectable } from '@angular/core';
import skillsData from '../../resources/sfia-data/skills.json';
import skillCriteriaData from '../../resources/sfia-data/sfia-criteria.json';

@Injectable({
  providedIn: 'root'
})
export class SfiaSkillsService {

  constructor() { }

  getSFIASkills(): any {
    return skillsData;
  }

  getSFIASkillsCriteria(): any[] {
    return skillCriteriaData;
  }

  getSFIASkillCriiteraByCodeAndLevel(skillCode: string, sfiaLevel: number) {
    let criteria = this.getSFIASkillsCriteria().filter(skill=> skill.skillCode == skillCode && skill.sfiaLevel == sfiaLevel);
    return criteria;
  }
}
