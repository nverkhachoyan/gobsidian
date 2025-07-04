---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/11
- monster/environment/mountain
- monster/size/huge
- monster/type/giant/cloud-giant
aliases: ["Cloud Giant Smiling One"]
NoteIcon: monster
BestiaryType: giant (cloud giant)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 146
---
# [Cloud Giant Smiling One](3-Mechanics\CLI\bestiary\giant/cloud-giant-smiling-one-vgm.md)
*Source: Volo's Guide to Monsters p. 146*  

Cloud giants aren't, on the whole, religious. They tolerate many conflicting ideas about their patron deity, Memnor. The smiling ones strain that tolerance.

Smiling ones are cloud giants who honor and emulate Memnor's craftiness and deceit above all else. They are tricksters supreme who use sleight of hand, deception, misdirection, and magic in their pursuit of wealth. They also possess a flair for unpredictability and a wicked sense of humor.

While cloud giants expect a certain amount of trickery and deceit in their dealings with others of their kind, smiling ones overstep the bounds of decorum with their behavior, doing and saying things that nobler cloud giants consider beneath the dignity of their kind.

## Mysterious Masks

Smiling ones take their name from the strange two-faced masks they wear. The smiling half of the face often looks more like a smirk or a triumphant sneer than a pleasant grin. The frowning half represents the displeasure smiling ones feel about their place in the ordning-second to the storm giants. The masks serve as symbols of their devotion, but they also conceal their wearers' true facial expressions.

```statblock
"name": "Cloud Giant Smiling One (VGM)"
"size": "Huge"
"type": "giant"
"subtype": "cloud giant"
"alignment": "Chaotic Neutral"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "262"
"hit_dice": "21d12 + 128"
"stats":
- !!int "26"
- !!int "12"
- !!int "22"
- !!int "15"
- !!int "16"
- !!int "17"
"speed": "40 ft."
"saves":
  "Wisdom": !!int "7"
  "Intelligence": !!int "6"
  "Constitution": !!int "10"
"skillsaves":
  "Sleight of Hand": !!int "9"
  "Deception": !!int "11"
  "Insight": !!int "7"
  "Perception": !!int "7"
"senses": "passive Perception 17"
"languages": "Common, Giant"
"cr": "11"
"traits":
- "desc": "The giant's innate spellcasting ability is Charisma (spell save DC 15).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md),\
    \ [light](/3-Mechanics/CLI/spells/light.md)\n\n1/day each: [control weather](/3-Mechanics/CLI/spells/control-weather.md),\
    \ [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md)\n\n3/day each: [feather\
    \ fall](/3-Mechanics/CLI/spells/feather-fall.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md), [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Innate Spellcasting"
- "desc": "The giant is a 5th-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 15, dice: d20+7 (+7) to hit with spell attacks). The giant\
    \ has the following bard spells prepared:\n\nCantrips (at will): [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md), [vicious mockery](/3-Mechanics/CLI/spells/vicious-mockery.md)\n\
    \n1st level (4 slots): [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [silent image](/3-Mechanics/CLI/spells/silent-image.md),\
    \ [Tasha's hideous laughter](/3-Mechanics/CLI/spells/tashas-hideous-laughter.md)\n\
    \n2nd level (3 slots): [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n3rd level (2 slots):\
    \ [major image](/3-Mechanics/CLI/spells/major-image.md), [tongues](/3-Mechanics/CLI/spells/tongues.md)"
  "name": "Spellcasting"
- "desc": "The giant has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
"actions":
- "desc": "The giant makes two attacks with its morningstar."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d8 + 8|text(21) (3d8 + 8) bludgeoning damage. The attack\
    \ deals an extra dice:4d6|text(14) (4d6) damage if the giant has advantage\
    \ on the attack roll."
  "name": "Morningstar"
- "desc": "Ranged Weapon Attack: dice: d20+12 (+12) to hit, range 60/240 ft.,\
    \ one target. Hit: dice:4d10 + 8|text(30) (4d10 + 8) bludgeoning damage.\
    \ The attack deals an extra dice:4d6|text(14) (4d6) damage if the giant has\
    \ advantage on the attack roll."
  "name": "Rock"
- "desc": "The giant magically polymorphs into a beast or humanoid it has seen, or\
    \ back into its true form. Any equipment the giant is wearing or carrying is absorbed\
    \ by the new form. Its statistics, other than its size, are the same in each form.\
    \ It reverts to its true form if it dies."
  "name": "Change Shape"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Cloud%20Giant%20Smiling%20One.webp"
```
^statblock

```encounter-table
name: Cloud Giant Smiling One
creatures:
 - 1: Cloud Giant Smiling One
```

## Environment

mountain