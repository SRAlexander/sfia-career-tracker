import { Injectable } from '@angular/core';
import data from '../../resources/roles-data/roles.json'

@Injectable({
  providedIn: 'root'
})
export class RolesService {

  constructor() { }

  getFullRoles(): any[] {
    return JSON.parse(JSON.stringify(data));
  }

  getRoleTypes(): any[] {
    var primaryRoles = [];
    primaryRoles = this.getFullRoles().map(roleData => roleData.type);
    return primaryRoles;
  }

  getFullRoleBySFIAAndType(sfiaLevel: number, roleType: string) {
    let response = this.getFullRoles().filter(role => role.sfiaLevel === sfiaLevel && role.type === roleType)
    return response;
  }

}
