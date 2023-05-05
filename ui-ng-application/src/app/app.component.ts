import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})


export class AppComponent {

  title = 'ui-ng-application';

  selectedPDP: any = null;
  stage: number = 0;


  onPDPSelected(pdp: any) {
    this.selectedPDP = pdp;
  }

  onReturnClick() {
    this.selectedPDP = null;
  }

  stageChange() {
    this.stage = 1;
  }
}



