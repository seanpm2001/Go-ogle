Sex ::= male | female.

person <~ {dayOfBirth:[] => day.
           age:[] => integer.
           sex:[] => Sex.
           name:[] => string.
           home:[] => string.
           lives:[string]{}}.
  
person:[string, day, Sex, string] $= person.
  
person(Nm, Born, Sx, Hm)..{
  dayOfBirth() => Born.
  age() => yearsBetween(now(), Born).
  sex() => Sx.
  name() => Nm.
  home() => Hm.
  lives(Pl) :- Pl = home().
  yearsBetween:[integer, day] => integer.
  yearsBetween(...) => ..
}.
  
newPerson:[string, day, Sex, string] => person.
  
newPerson(Nm, Born, Sx, Hm) => $person(Nm, Born, Sx, Hm).
