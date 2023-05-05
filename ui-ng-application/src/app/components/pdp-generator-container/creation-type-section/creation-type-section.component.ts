import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { throwError } from 'rxjs';

@Component({
  selector: 'app-creation-type-section',
  templateUrl: './creation-type-section.component.html',
  styleUrls: ['./creation-type-section.component.scss']
})
export class CreationTypeSectionComponent implements OnInit {

  @Input() sfiaLevel : number = 0;
  @Input() creationType : string = "";
  @Output() creationTypeSelected : EventEmitter<string> = new EventEmitter();
  
  roleSelected: boolean = false;
  skillSelected: boolean = false; 

  ngOnInit() {
    if (this.creationType == "skills") {
      this.skillSelected = true;
    }

    if (this.creationType == "roles") {
      this.roleSelected = true;
    }
  }

  onRoleSelected() {
    this.roleSelected = true;
    this.skillSelected = false;
    this.creationTypeSelected.emit("roles")
  }

  onSkillsSelected() {
    this.roleSelected = false;
    this.skillSelected = true;
    this.creationTypeSelected.emit("skills")
  }
}
