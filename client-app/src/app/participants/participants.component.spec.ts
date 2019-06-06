import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ParticipantsComponentComponent } from './participants.component';

describe('ParticipantsComponentComponent', () => {
  let component: ParticipantsComponentComponent;
  let fixture: ComponentFixture<ParticipantsComponentComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ParticipantsComponentComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ParticipantsComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
