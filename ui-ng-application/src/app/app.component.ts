import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'ui-ng-application';

  selectedPDP: any = null;

  onPDPSelected(pdp: any) {
    this.selectedPDP = pdp;
  }
}



