---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/environment/underdark
- monster/size/small
- monster/type/humanoid/derro
aliases: ["Derro Savant"]
NoteIcon: monster
BestiaryType: humanoid (derro)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 159
---
# [Derro Savant](3-Mechanics\CLI\bestiary\humanoid/derro-savant-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 159*  

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
"name": "Derro Savant (MTF)"
"size": "Small"
"type": "humanoid"
"subtype": "derro"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "36"
"hit_dice": "8d6 + 8"
"stats":
- !!int "9"
- !!int "14"
- !!int "12"
- !!int "11"
- !!int "5"
- !!int "14"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
"senses": "darkvision 120 ft., passive Perception 7"
"languages": "Dwarvish, Undercommon"
"cr": "3"
"traits":
- "desc": "The derro savant is a 5th-level spellcaster. Its spellcasting ability is\
    \ Charisma (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks).\
    \ The derro knows the following sorcerer spells:\n\nCantrips (at will): [acid\
    \ splash](/3-Mechanics/CLI/spells/acid-splash.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md),\
    \ [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md)\n\n1st level (4 slots):\
    \ [burning hands](/3-Mechanics/CLI/spells/burning-hands.md), [chromatic orb](/3-Mechanics/CLI/spells/chromatic-orb.md),\
    \ [sleep](/3-Mechanics/CLI/spells/sleep.md)\n\n2nd level (3 slots): [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [spider climb](/3-Mechanics/CLI/spells/spider-climb.md)\n\n3rd level (2 slots):\
    \ [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)"
  "name": "Spellcasting"
- "desc": "The derro savant has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "While in sunlight, the derro savant has disadvantage on attack rolls, as\
    \ well as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage."
  "name": "Quarterstaff"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Derro%20Savant.webp"
```
^statblock

```encounter-table
name: Derro Savant
creatures:
 - 1: Derro Savant
```

## Environment

underdark