---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/forest
- monster/size/large
- monster/type/monstrosity
aliases: ["Displacer Beast"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 81
---
# [Displacer Beast](3-Mechanics\CLI\bestiary\monstrosity/displacer-beast.md)
*Source: Monster Manual p. 81*  

This monstrous predator takes its name from its ability to displace light so that it appears to be several feet away from its actual location. A displacer beast resembles a sleek great cat covered in blue-black fur.

However, its otherworldly origins are clear in its six legs and the two tentacles sprouting from its shoulders, both ending in pads tipped with spiky protrusions. A displacer beast's eyes glow with an awful malevolence that persists even in death.

## Unseelie Origins

Displacer beasts roamed the twilight lands of the Feywild for ages, until they were captured and trained by the Unseelie Court. The warriors of the court selectively bred the beasts to reinforce their ferocious and predatory nature, using them to hunt unicorns, pegasi, and other wondrous prey. However, it didn't take long for the displacer beasts to use their malevolent intelligence to escape their masters.

Running and breeding freely in the Feywild, the displacer beasts soon came to the attention of the Seelie Court. With blink dog companions at their side, fey hunters drove these predators to the fringes of the Feywild, where many crossed over to the Material Plane. To this day, displacer beasts and blink dogs attack each other on sight.

## Love of the Kill

Displacer beasts kill not just for food but also for sport. They target prey even when not hungry, often toying with their victims to entertain themselves until they are ready to eat. After killing its prey using its tentacles, a displacer beast drags the corpse to a quiet place where it can feed without distraction.

Displacer beasts hunt alone or in small prides that demonstrate skill at setting ambushes. A single beast will strike and withdraw, luring prey into a densely wooded area where its pack mates wait. Packs of displacer beasts hunting near trade roads recall the frequency and schedule of regular caravans, laying down ambushes to pick off those caravans.

```statblock
"name": "Displacer Beast"
"size": "Large"
"type": "monstrosity"
"alignment": "Lawful Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "85"
"hit_dice": "10d10 + 30"
"stats":
- !!int "18"
- !!int "15"
- !!int "16"
- !!int "6"
- !!int "12"
- !!int "8"
"speed": "40 ft."
"senses": "darkvision 60 ft., passive Perception 11"
"languages": ""
"cr": "3"
"traits":
- "desc": "If the displacer beast is subjected to an effect that allows it to make\
    \ a saving throw to take only half damage, it instead takes no damage if it succeeds\
    \ on the saving throw, and only half damage if it fails."
  "name": "Avoidance"
- "desc": "The displacer beast projects a magical illusion that makes it appear to\
    \ be standing near its actual location, causing attack rolls against it to have\
    \ disadvantage. If it is hit by an attack, this trait is disrupted until the end\
    \ of its next turn. This trait is also disrupted while the displacer beast is\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated) or has a\
    \ speed of 0."
  "name": "Displacement"
"actions":
- "desc": "The displacer beast makes two attacks with its tentacles."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) bludgeoning damage plus dice:1d6|text(3)\
    \ (1d6) piercing damage."
  "name": "Tentacle"
"source":
- "MM"
- "TftYP"
- "WDH"
- "WDMM"
- "ERLW"
- "WBtW"
- "JttRC"
- "KftGV"
- "SatO"
- "BMT"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Displacer%20Beast.webp"
```
^statblock

```encounter-table
name: Displacer Beast
creatures:
 - 1: Displacer Beast
```

## Environment

forest