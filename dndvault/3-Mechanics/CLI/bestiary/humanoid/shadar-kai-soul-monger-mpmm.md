---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Shadar-kai Soul Monger"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 214, Mordenkainen's Tome of Foes p. 226
---
# [Shadar-kai Soul Monger](3-Mechanics\CLI\bestiary\humanoid/shadar-kai-soul-monger-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 214, Mordenkainen's Tome of Foes p. 226*  

```statblock
"name": "Shadar-kai Soul Monger (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Typically  Neutral Evil"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "136"
"hit_dice": "21d8 + 42"
"stats":
- !!int "8"
- !!int "17"
- !!int "14"
- !!int "19"
- !!int "16"
- !!int "13"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Dexterity": !!int "7"
  "Wisdom": !!int "7"
"skillsaves":
  "Perception": !!int "7"
"damage_immunities": "necrotic, psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 17"
"languages": "Common, Elvish"
"cr": "11"
"traits":
- "desc": "The shadar-kai casts one of the following spells, requiring no material\
    \ components and using Intelligence as the spellcasting ability (spell save DC\
    \ 16):\n\n1/day each: [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md),\
    \ [finger of death](/3-Mechanics/CLI/spells/finger-of-death.md), [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md),\
    \ [seeming](/3-Mechanics/CLI/spells/seeming.md)"
  "name": "Spellcasting"
- "desc": "The shadar-kai has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
- "desc": "The shadar-kai has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "When it reduces a creature to 0 hit points, the shadar-kai can gain temporary\
    \ hit points equal to half the creature's hit point maximum. While the shadar-kai\
    \ has temporary hit points from this trait, it has advantage on attack rolls."
  "name": "Soul Thirst"
- "desc": "Any Beast or Humanoid (except an elf) that starts its turn within 5 feet\
    \ of the shadar-kai has its speed reduced by 20 feet until the start of that creature's\
    \ next turn."
  "name": "Weight of Ages"
"actions":
- "desc": "The shadar-kai makes two Shadow Dagger attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:4d4 + 3|text(13) (4d4 + 3) piercing\
    \ damage plus dice:3d12|text(19) (3d12) necrotic damage, and the target has\
    \ disadvantage on saving throws until the end of the shadar-kai's next turn. Hit\
    \ or Miss: The dagger magically returns to the shadar-kai's hand immediately after\
    \ a ranged attack."
  "name": "Shadow Dagger"
- "desc": "The shadar-kai emits weariness in a 60-foot cube. Each creature in that\
    \ area must make a DC 16 Constitution saving throw. On a failed save, a creature\
    \ takes dice:10d8|text(45) (10d8) psychic damage and suffers 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion).\
    \ On a successful save, it takes half as much damage and doesn't gain a level\
    \ of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion)."
  "name": "Wave of Weariness (Recharge 4-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Shadar-kai%20Soul%20Monger.webp"
```
^statblock

```encounter-table
name: Shadar-kai Soul Monger
creatures:
 - 1: Shadar-kai Soul Monger
```

## Environment

underdark, urban