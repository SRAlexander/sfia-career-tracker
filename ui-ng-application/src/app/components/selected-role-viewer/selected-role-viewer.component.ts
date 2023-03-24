import { Component, Input, OnInit } from '@angular/core';
import { SfiaSkillsService } from 'src/app/services/sfia-skills.service';

@Component({
  selector: 'app-selected-role-viewer',
  templateUrl: './selected-role-viewer.component.html',
  styleUrls: ['./selected-role-viewer.component.scss']
})
export class SelectedRoleViewerComponent implements OnInit{

  @Input() roleDefinition: any;

  skills: any[] = [];

  constructor(private _skillService: SfiaSkillsService){

  }

  ngOnInit(): void {
    this.skills = this._skillService.getSFIASkills();
  }

  getSkillFullNameFromCode(code: string) {
    var foundskills = this.skills.filter(skill => skill.code === code)
    if (foundskills.length > 0){
      return foundskills[0].title;
    } else {
      return "Unknown Skill"
    }
  }
}
