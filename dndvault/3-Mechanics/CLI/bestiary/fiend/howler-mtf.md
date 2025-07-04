---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/8
- monster/environment/desert
- monster/environment/grassland
- monster/environment/hill
- monster/environment/underdark
- monster/size/large
- monster/type/fiend
aliases: ["Howler"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 210
---
# [Howler](3-Mechanics\CLI\bestiary\fiend/howler-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 210*  

> [!quote]- A quote from Mordenkainen  
> 
> Why does the howler sing? Doing so causes its prey to flee, and surely stealth would make for better hunting in howling Pandemonium. There is only one answer: the creature can taste fear.

## Howler

A far-off wail precedes the sight of a howler. Even at a distance, one's mind cringes at the sound and fills with horror at the realization that the noise is drawing closer. When howlers go on the prowl, courage isn't enough to stand up against them, and even one's sanity is at risk.

## Prowlers from Pandemonium

These nightmare creatures, native to Pandemonium, can also be found on most of the Lower Planes, because of the many fiends that capture them and train them as war hounds. Howlers can be domesticated, after a fashion, but they respond only to brutal training during which they are forced to recognize the trainer as the pack's undisputed leader. A trained pack then follows its leader without hesitation. Howler packs course over the battlefields of the Blood War and also serve evil mortals who have the power and the savagery to command their loyalty.

## Brutal Hunters

Howlers rely on speed, numbers, and their mind-numbing howl to corner prey before they tear it apart. The howl floods the minds of its victims, making complex thought impossible. They can do little more than stare in horror and stumble around the battlefield in a search for safety. Any task more demanding than that is beyond their capability. Fiends especially prize howlers for this reason, because for a few crucial moments in a battle, their howls can neutralize an enemy's ability to use spells and other powers.

```statblock
"name": "Howler (MTF)"
"size": "Large"
"type": "fiend"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "90"
"hit_dice": "12d10 + 24"
"stats":
- !!int "17"
- !!int "16"
- !!int "15"
- !!int "5"
- !!int "20"
- !!int "6"
"speed": "40 ft."
"skillsaves":
  "Perception": !!int "8"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "understands Abyssal but can't speak"
"cr": "8"
"traits":
- "desc": "A howler has advantage on attack rolls against a creature if at least one\
    \ of the howler's allies is within 5 feet of the creature and the ally isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
"actions":
- "desc": "The howler makes two bite attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage. If the target is\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) it takes an additional\
    \ dice:4d10|text(22) (4d10) psychic damage. This attack ignores damage resistance."
  "name": "Rending Bite"
- "desc": "The howler emits a keening howl in a 60-foot cone. Each creature in that\
    \ area that isn't [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened) must\
    \ succeed on a DC 16 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the end of the howler's next turn. While a creature is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, its speed is halved, and it is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\
    \ A target that successfully saves is immune to the Mind-Breaking Howl of all\
    \ howlers for the next 24 hours."
  "name": "Mind-Breaking Howl (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Howler.webp"
```
^statblock

```encounter-table
name: Howler
creatures:
 - 1: Howler
```

## Environment

desert, grassland, hill, underdark