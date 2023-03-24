import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { RolesService } from 'src/app/services/roles.service';

@Component({
  selector: 'app-role-selector',
  templateUrl: './role-selector.component.html',
  styleUrls: ['./role-selector.component.scss']
})
export class RoleSelectorComponent implements OnInit {

  @Input() sfiaLevel : number = 3;
  @Output() onRoleSelected = new EventEmitter<any>();

  primaryRoles: any[] = [];

  constructor(private _rolesService: RolesService){

  }

  ngOnInit(): void {
    this.primaryRoles = [...new Set(this._rolesService.getRoleTypes())];
    this.primaryRoles.splice(0, 0, "Select a role type...")
  }

  onRoleSelectedEvent(event: any) {
    var roleSearchResult = this._rolesService.getFullRoleBySFIAAndType(this.sfiaLevel, event.target.value)
    if (roleSearchResult.length > 0){
      this.onRoleSelected.emit(roleSearchResult[0]);
    } else {
      this.onRoleSelected.emit(null);
    }
  }
}
