import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { PdpGeneratorService } from 'src/app/services/pdp-generator.service';

@Component({
  selector: 'app-sfia-selector-form',
  templateUrl: './sfia-selector-form.component.html',
  styleUrls: ['./sfia-selector-form.component.scss']
})
export class SfiaSelectorFormComponent implements OnInit{
  
  constructor(private _pdpGeneratorService : PdpGeneratorService) {

  }

  @Output() pdpSelected: EventEmitter<any> = new EventEmitter();

  roleSelectorEnabled: Boolean = true;
  skillSelectorEnabled: Boolean = false;
  
  sfiaOptions : any[] = [ { "level" : 0, "display" : "Select a SFIA level..."}, 
                        { "level" : 1, "display" : "SFIA 1: Graduate"}, 
                        { "level" : 2, "display" : "SFIA 2: Level 1"},
                        { "level" : 3, "display" : "SFIA 3: Level 2"},
                        { "level" : 4, "display" : "SFIA 4: Senior"},
                        { "level" : 5, "display" : "SFIA 5: Lead"},
                        { "level" : 6, "display" : "SFIA 6: Prinicipal"}];

  selectedSFIALevel : number = 0;
  selectedRole: any = null;

  ngOnInit(): void {
    throw new Error('Method not implemented.');
  }

  onRoleSelected() {
    this.roleSelectorEnabled = true;
    this.skillSelectorEnabled = false;
  }

  onSkillsSelected() {
    this.roleSelectorEnabled = false;
    this.skillSelectorEnabled = true;
    this.selectedRole = null;
  }

  onSFIASelected(event: any) {
    var selectedSfiaText = event.target.value;
    this.selectedSFIALevel = this.sfiaOptions.filter(sfiaOption => sfiaOption.display === selectedSfiaText)[0].level;
    console.log(this.selectedSFIALevel);
  }

  roleSelected(selectedRole : any) {
    this.selectedRole = selectedRole;
  }

  generatePDP(){
    var pdpJson = this._pdpGeneratorService.generatePDP(this.selectedRole)
    this.pdpSelected.emit(pdpJson);
  }

}
