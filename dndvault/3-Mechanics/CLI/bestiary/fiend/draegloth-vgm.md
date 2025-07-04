---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/7
- monster/environment/underdark
- monster/size/large
- monster/type/fiend/demon
aliases: ["Draegloth"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 141
---
# [Draegloth](3-Mechanics\CLI\bestiary\fiend/draegloth-vgm.md)
*Source: Volo's Guide to Monsters p. 141*  

A draegloth is a half-drow, half-glabrezu demon, born of a drow high priestess in an unholy, dangerous ritual. Gifted with innate magic and physical might, it usually remains in the service of its mother's house, lending its thirst for destruction to that house's plans to triumph over its rivals.

A draegloth is an ogre-sized, four-armed humanoid with purple-black skin and yellow-white hair. Two of its arms are huge and muscular, tipped with sharp claws; the other two are the size and shape of drow arms, capable of delicate movements. Although the creature is heavily muscled, it is graceful and quiet like a drow. Its face is clearly demonic, with bestial features, glowing red eyes, an elongated dog-like snout, and a mouth full of sharp teeth.

## Blessing on the House

The ritual to create a draegloth succeeds only rarely, but when it does, it is a great event that is seen by the drow of the house as a sign of the demon lord Lolth's favor-and a sign of Lolth's disregard for the family's rivals, which were not thus gifted. The birth prompts the leaders of the house to begin crafting new plans to strike at its rivals when the draegloth is fully grown. These plans always use the draegloth in a significant role, because its abilities can turn the tide in a battle against a house that doesn't have a draegloth of its own.

These drow house pets are as graceful and nimble as Waterdhavian stage dancers. Only they're slayers and enforcers, four-armed brutes built like an ogre. Life isn't fair.

-Volo

## Subservient Enforcers

Although it plays an important part in the welfare of its house, a draegloth can't rise above the status of a favored slave or a consort to a priestess. Before a draegloth is given any duties, it receives instruction in accepting the role set for it and not challenging authority. Draegloths instinctively resist this sort of treatment, but most of them take out their frustration on their house's enemies. A draegloth that can't suppress its ambitions might abandon its house and strike out on its own. Whether these rebellious draegloths are part of Lolth's plan for sowing even greater chaos is unclear.

## Brute Cunning and Dark Magic

A draegloth loves the feeling of tearing opponents apart with its claws and teeth and of wielding the magic that courses through its veins. Most are too impatient to bother with complicated tactics, but a few go on to learn more destructive magic.

```statblock
"name": "Draegloth (VGM)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "123"
"hit_dice": "13d10 + 52"
"stats":
- !!int "20"
- !!int "15"
- !!int "18"
- !!int "13"
- !!int "11"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "3"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 13"
"languages": "Abyssal, Elvish, Undercommon"
"cr": "7"
"traits":
- "desc": "The draegloth's innate spellcasting ability is Charisma (spell save DC\
    \ 11). The draegloth can innately cast the following spells, requiring no material\
    \ components:\n\nAt will: [darkness](/3-Mechanics/CLI/spells/darkness.md)\n\
    \n1/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md), [dancing\
    \ lights](/3-Mechanics/CLI/spells/dancing-lights.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md)"
  "name": "Innate Spellcasting"
- "desc": "The draegloth has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
"actions":
- "desc": "The draegloth makes three attacks: one with its bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d10 + 5|text(16) (2d10 + 5) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d10 + 5|text(16) (2d10 + 5) slashing damage."
  "name": "Claws"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Draegloth.webp"
```
^statblock

```encounter-table
name: Draegloth
creatures:
 - 1: Draegloth
```

## Environment

underdark