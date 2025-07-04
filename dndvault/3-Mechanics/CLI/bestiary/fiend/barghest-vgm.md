---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/fiend/shapechanger
aliases: ["Barghest"]
NoteIcon: monster
BestiaryType: fiend (shapechanger)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 123
---
# [Barghest](3-Mechanics\CLI\bestiary\fiend/barghest-vgm.md)
*Source: Volo's Guide to Monsters p. 123*  

Long ago, Maglubiyet, master of the goblinoid gods, bargained with the General of Gehenna for aid. The General provided yugoloths that died to serve the cause of the goblin god. Yet when the time came to honor his part of the compact, Maglubiyet reneged on the deal. As an act of vengeance, the General of Gehenna created the soul-devouring barghests to devour goblinoid souls and deprive Maglubiyet of troops for his army in the afterlife.

## Consumers of Souls

A barghest is born to goblin parents just as normal offspring are. The creature emerges in the form of a goblin, then develops the ability to assume its true form: that of a large, fiendish canine.

The mission of every barghest, implanted in it by the General of Gehenna, is to consume seventeen goblinoid souls by devouring the bodies of those it kills. Souls consumed in this way are prevented from joining Maglubiyet's forces in Acheron. Why seventeen? Because the oaths Maglubiyet broke in his compact with the General totaled seventeen.

A barghest hungers for the day when it can complete its mission, return to Gehenna, and serve the General directly in his yugoloth legions, but it doesn't kill goblinoids indiscriminately. By devouring the souls of goblinoid leaders and other powerful individuals, rather than lowly goblins, a barghest earns elevated status in the afterlife. Barghests typically keep their true nature secret, preying upon a goblin or two when the opportunity arises, until they reach adult age and are old and strong enough to seek out stronger prey. When goblins discover that a barghest is among them, they react with groveling obeisance, each member of the tribe eager to show the barghest that it isn't worthy of being devoured.

## Banished by Fire

A barghest avoids contact with large, open fires. Any conflagration larger than its body acts as a gateway to Gehenna and banishes the fiend to that plane, where it is likely to be slain or enslaved by a yugoloth for its failure.

## Soul Feeding

A barghest can feed on the corpse of a humanoid that it killed that has been dead for less than 10 minutes, devouring both flesh and soul in doing so. This feeding takes at least 1 minute, and it destroys the victim's body. The victim's soul is trapped in the barghest for 24 hours, after which time it is digested. If the barghest dies before the soul is digested, the soul is released.

While a humanoid's soul is trapped in a barghest, any form of revival that could work has only a 50% chance chance of doing so, freeing the soul from the barghest if it is successful. Once a creature's soul is digested, however, no mortal magic can return that humanoid to life.

```statblock
"name": "Barghest (VGM)"
"size": "Large"
"type": "fiend"
"subtype": "shapechanger"
"alignment": "Neutral Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "90"
"hit_dice": "12d10 + 24"
"stats":
- !!int "19"
- !!int "15"
- !!int "14"
- !!int "13"
- !!int "12"
- !!int "14"
"speed": "60 ft. (30 ft. in goblin form)"
"skillsaves":
  "Intimidation": !!int "4"
  "Deception": !!int "4"
  "Stealth": !!int "4"
  "Perception": !!int "5"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 15"
"languages": "Abyssal, Common, Goblin, Infernal, telepathy 60 ft."
"cr": "4"
"traits":
- "desc": "The barghest's innate spellcasting ability is Charisma (spell save DC 12).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [levitate](/3-Mechanics/CLI/spells/levitate.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [pass without trace](/3-Mechanics/CLI/spells/pass-without-trace.md)\n\n1/day\
    \ each: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Innate Spellcasting"
- "desc": "The barghest can use its action to polymorph into a Small goblin or back\
    \ into its true form. Other than its size and speed, its statistics are the same\
    \ in each form. Any equipment it is wearing or carrying isn't transformed. The\
    \ barghest reverts to its true form if it dies."
  "name": "Shapechanger"
- "desc": "When the barghest starts its turn engulfed in flames that are at least\
    \ 10 feet high or wide, it must succeed on a DC 15 Charisma saving throw or be\
    \ instantly banished to Gehenna. Instantaneous bursts of flame (such as a red\
    \ dragon's breath or a [fireball](/3-Mechanics/CLI/spells/fireball.md) spell)\
    \ don't have this effect on the barghest."
  "name": "Fire Banishment"
- "desc": "The barghest has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
- "desc": "A barghest can feed on the corpse of a humanoid that it killed that has\
    \ been dead for less than 10 minutes, devouring both flesh and soul in doing so.\
    \ This feeding takes at least 1 minute, and it destroys the victim's body. The\
    \ victim's soul is trapped in the barghest for 24 hours, after which time it is\
    \ digested. If the barghest dies before the soul is digested, the soul is released.\n\
    \nWhile a humanoid's soul is trapped in a barghest, any form of revival that could\
    \ work has only a 50% chance chance of doing so, freeing the soul from the barghest\
    \ if it is successful. Once a creature's soul is digested, however, no mortal\
    \ magic can return that humanoid to life."
  "name": "Soul Feeding"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage."
  "name": "Claws"
"source":
- "VGM"
- "TftYP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Barghest.webp"
```
^statblock

```encounter-table
name: Barghest
creatures:
 - 1: Barghest
```

## Environment

underdark, mountain, grassland, forest, hill