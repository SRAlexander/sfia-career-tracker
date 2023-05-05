import { Component, ElementRef, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { NgbTypeahead, NgbTypeaheadModule } from '@ng-bootstrap/ng-bootstrap';
import { Observable, Subject, merge, OperatorFunction } from 'rxjs';
import { debounceTime, distinctUntilChanged, filter, map } from 'rxjs/operators';
import { FormsModule } from '@angular/forms';
import { JsonPipe } from '@angular/common';


@Component({
  selector: 'app-custom-typeahead',
  templateUrl: './custom-typeahead.component.html',
  styleUrls: ['./custom-typeahead.component.scss'],
  standalone: true,
  imports: [NgbTypeaheadModule, FormsModule, JsonPipe],
})
export class CustomTypeaheadComponent {


  @Input() selectableList: any[] = [];
  @Output() itemSelected: EventEmitter<any> = new EventEmitter();

  model: any;

  @ViewChild('instance', { static: true })
  instance!: NgbTypeahead;
  focus$ = new Subject<string>();
  click$ = new Subject<string>();

  search: OperatorFunction<string, readonly string[]> = (text$: Observable<string>) => {
    const debouncedText$ = text$.pipe(debounceTime(200), distinctUntilChanged());
    const clicksWithClosedPopup$ = this.click$.pipe(filter(() => !this.instance.isPopupOpen()));
    const inputFocus$ = this.focus$;

    return merge(debouncedText$, inputFocus$, clicksWithClosedPopup$).pipe(
      map((term) =>
        (term === '' ? this.selectableList : this.selectableList.filter((v) => v.title.toLowerCase().indexOf(term.toLowerCase()) > -1)).slice(0, 100),
      ),
    );
  };

  formatListValue(value: any)   {
    if(value.title)
      return value.title
    return value;
  }

  onItemSelected($event: any) {
    this.itemSelected.emit($event.item)
    $event.preventDefault()
    this.model = "";
  }
}

