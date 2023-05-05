import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { PdpGeneratorService } from 'src/app/services/pdp-generator.service';

@Component({
  selector: 'app-sfia-selector-section',
  templateUrl: './sfia-selector-section.component.html',
  styleUrls: ['./sfia-selector-section.component.scss']
})
export class SfiaSelectorFormComponent implements OnInit{
  
  @Input() sfiaLevel : number = 0;
  @Output() sfiaSelected = new EventEmitter<number>

  ngOnInit(): void {
    this.selectedSFIA = this.sfiaLevel;
  }

  roleSelectorEnabled: Boolean = true;
  skillSelectorEnabled: Boolean = false;
  selectedSFIA: number = 0;

  sfiaOptions : any[] = [ { "level" : 0, "display" : "Select a SFIA level..."}, 
                        { "level" : 1, "display" : "SFIA 1: Graduate"}, 
                        { "level" : 2, "display" : "SFIA 2: Level 1"},
                        { "level" : 3, "display" : "SFIA 3: Level 2"},
                        { "level" : 4, "display" : "SFIA 4: Senior"},
                        { "level" : 5, "display" : "SFIA 5: Lead"},
                        { "level" : 6, "display" : "SFIA 6: Prinicipal"}];

  onSFIASelected(event: any) {
    this.sfiaSelected.emit(event.level);
  }

}
