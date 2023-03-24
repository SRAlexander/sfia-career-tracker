import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-pdp-viewer',
  templateUrl: './pdp-viewer.component.html',
  styleUrls: ['./pdp-viewer.component.scss']
})
export class PdpViewerComponent {

  @Input() PDPJson: any;

  
}
