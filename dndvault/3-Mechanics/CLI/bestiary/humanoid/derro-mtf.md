---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1-4
- monster/environment/mountain
- monster/environment/underdark
- monster/size/small
- monster/type/humanoid/derro
aliases: ["Derro"]
NoteIcon: monster
BestiaryType: humanoid (derro)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 158
---
# [Derro](3-Mechanics\CLI\bestiary\humanoid/derro-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 158*  

> [!quote]- A quote from Mordenkainen  
> 
> Mad as hatters, every one of them. No derro has ever been sane, but neither is every form of madness a hindrance.

## Derro

Derro slink through the subterranean realms, seeking places that are safe from the perils of the Underdark. Equal parts fearful and vicious, bands of these dwarfkin prey on those weaker than themselves, while giving simpering obeisance to any creatures they deem more powerful. Wild-haired, haggard, shuddering, and shabbily dressed, a lone derro seems a pitiable creature, but when a cackling, spitting, growling, howling horde of them attacks, the sight inspires both fear and revulsion.

## Madness and Sorcery

Fractious in groups and individually weak, derro would have been driven to extinction long ago but for two elements of their character. They have an inborn tendency toward paranoia, a peculiarity that serves them well as they navigate the dangers of the Underdark and its societies. They also have a stronger-than-normal tendency to develop sorcerous power. Individuals who do so become their leaders, known as savants. The derro consider these sorcerers to be specially blessed by their deity, Diirinka.

## Forgotten Duergar

Grandiose fantasies and rampant fanaticism have obscured the true origin of the derro, even among themselves. Most dwarves don't recognize derro as kin, but the legends that the derro tell about their race and the story that the duergar believe share a grain of truth.

According to the duergar, the derro are descended from dwarves of a clan that was left behind when the others escaped the mind flayers' rule. They eventually also got away, but not before becoming demented and contorted.

The derro tell their own story of flight and survival in the Underdark, and the mind flayers aren't always the enemy. Laduguer and Deep Duerra don't feature in their mythic history. Instead they tell of two brothers, Diirinka and Diinkarazan, and of how Diirinka cleverly betrayed his sibling so that he could steal magical power from the evil they escaped. The danger the brothers are said to face in this legend varies, depending on whatever foe the savants want to lead their people against, yet the essence of the story remains the same: a lesson of survival at any price and an example of how deceitfulness and cruelty can be virtues.

## Derro Madness

All derro suffer from a form of madness that most often manifests as mania and paranoia, but other mental afflictions and strange tics also commonly affect them. Derro take little notice of odd behavior in their ranks, except when an individual displays the characteristics of a savant. They believe the strange behavior of savants arises because those leaders carry messages from Diirinka. You can use the [Derro Madness](/3-Mechanics/CLI/tables/derro-madness-mtf.md) table to generate one or more odd qualities for a derro NPC.

```statblock
"name": "Derro (MTF)"
"size": "Small"
"type": "humanoid"
"subtype": "derro"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "13"
"hit_dice": "3d6 + 3"
"stats":
- !!int "10"
- !!int "14"
- !!int "12"
- !!int "11"
- !!int "5"
- !!int "9"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
"senses": "darkvision 120 ft., passive Perception 7"
"languages": "Dwarvish, Undercommon"
"cr": "1/4"
"traits":
- "desc": "The derro has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "While in sunlight, the derro has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6|text(3) (1d6) piercing damage. If the target is Medium or\
    \ smaller, the derro can choose to deal no damage and knock it [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Hooked Spear"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Light Crossbow"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Derro.webp"
```
^statblock

```encounter-table
name: Derro
creatures:
 - 1: Derro
```

## Environment

mountain, underdark