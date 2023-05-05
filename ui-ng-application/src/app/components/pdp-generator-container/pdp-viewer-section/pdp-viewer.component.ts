import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-pdp-viewer',
  templateUrl: './pdp-viewer.component.html',
  styleUrls: ['./pdp-viewer.component.scss']
})
export class PdpViewerComponent {

  @Input() pdpJson: any;
  @Output() returnEvent : EventEmitter<any> = new EventEmitter();

  returnClick() {
    this.returnEvent.emit(true);
  }

  downloadMarkdown(){

  }

  downloadCSV(){
    this.downloadCSVFile(this.pdpJson);
  }

  downloadCSVFile(data: any) {
    var flattenedGroups: any[] = []
    data.coreSkills.forEach((coreSkillGroup : any) => { coreSkillGroup.skillsGroup.forEach((sfiaCriteria:any) => {flattenedGroups.push(sfiaCriteria)}) })
    
    var header = Object.keys(flattenedGroups[0]);
    const csv = flattenedGroups.map((row : any) =>
      header
        .map((fieldName) => JSON.stringify(row[fieldName]))
        .join(',')
    );
    csv.unshift(header.join(','));
    const csvArray = csv.join('\r\n');
  
    const a = document.createElement('a');
    const blob = new Blob([csvArray], { type: 'text/csv' });
    const url = window.URL.createObjectURL(blob);
  
    a.href = url;
    const filename = Date.now().toString() + "_" + data.roleTitle.toLowerCase().replace(" ", "_") + ".csv"
    a.download = filename;
    a.click();
    window.URL.revokeObjectURL(url);
    a.remove();
  }

}
