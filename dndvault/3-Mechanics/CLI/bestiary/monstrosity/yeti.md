---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/arctic
- monster/size/large
- monster/type/monstrosity
aliases: ["Yeti"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 305. Available in the Basic Rules.
---
# [Yeti](3-Mechanics\CLI\bestiary\monstrosity/yeti.md)
*Source: Monster Manual p. 305. Available in the Basic Rules.*  

A yeti's windborne howl sounds out across remote mountains, striking fear into the hearts of the scattered miners and herders that dwell there. These hulking creatures stalk alpine peaks in a ceaseless hunt for food. Their snow-white fur lets them move like ghosts against the frozen landscape. A yeti's icy simian eyes can freeze its prey in place.

## Keen Hunters

Folk of the high peaks travel in groups and go armed, knowing that yetis can smell living flesh from miles away. When it finds prey, a yeti moves quickly over ice and stone to claim its meal, howling to the thrill of the hunt. Even in a blizzard, the scent of its quarry draws the yeti through the cold and snow.

Yetis hunt in solitude or in small family groups. When creatures flee from a yeti or engage it in battle, other yetis might catch the scent of blood and close in. The territorial yetis fight one another for the spoils of such battles, and yetis slain in the fight are also eaten, amid euphoric howls.

## Terrifying Howlers

Before an avalanche, a blizzard, or a deadly frost, the yetis' howls sweep down the mountain slopes on the icy wind. Some people of the alpine peaks believe that the voices of loved ones killed in avalanches and blizzards sound out in the wails of the yetis, crying warnings of ill omen. More pragmatic folk attest that the yeti's howl is a reminder that, despite the great accomplishments of civilization, the civilized become the hunted in nature's untamed domain.

## Brutal Rampagers

When mountain herds are abundant, yetis stay clear of humanoid realms. Driven by hunger, they attack humanoid settlements in waves, breaking down gates and stockade walls that once might have daunted them, then devouring the creatures within.

Devious mountain folk sometimes use the yetis as unwitting weapons. A warlord might lay down slaughtered sheep or goats to draw yetis into an enemy's camp, sowing chaos and thinning the ranks before battle. Mountain clan chiefs, wanting to expand their territory, overhunt local game to diminish the yetis' food supplies, inspiring attacks on humanoid settlements that are swiftly annexed in the aftermath.

## Abominable Yetis

An abominable yeti is larger than a normal yeti, standing three times as tall as a human. It typically lives and hunts alone, though a pair of abominable yetis might live together long enough to raise young. These towering yetis are highly territorial and savage, attacking and devouring any warm-blooded creatures they encounter, then scattering the bones across the ice and snow.

> [!quote]- A quote from Kelesta Hawke of the Emerald Enclave  
> 
> On your guard! That's not the wind howling!


```statblock
"name": "Yeti"
"size": "Large"
"type": "monstrosity"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "51"
"hit_dice": "6d10 + 18"
"stats":
- !!int "18"
- !!int "13"
- !!int "16"
- !!int "8"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 40 ft."
"skillsaves":
  "Stealth": !!int "3"
  "Perception": !!int "3"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Yeti"
"cr": "3"
"traits":
- "desc": "If the yeti takes fire damage, it has disadvantage on attack rolls and\
    \ ability checks until the end of its next turn."
  "name": "Fear of Fire"
- "desc": "The yeti has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
- "desc": "The yeti has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks made to hide in snowy terrain."
  "name": "Snow Camouflage"
"actions":
- "desc": "The yeti can use its Chilling Gaze and makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage plus dice:1d6|text(3)\
    \ (1d6) cold damage."
  "name": "Claw"
- "desc": "The yeti targets one creature it can see within 30 feet of it. If the target\
    \ can see the yeti, the target must succeed on a DC 13 Constitution saving throw\
    \ against this magic or take dice:3d6|text(10) (3d6) cold damage and then\
    \ be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed) for 1 minute,\
    \ unless it is immune to cold damage. The target can repeat the saving throw at\
    \ the end of each of its turns, ending the effect on itself on a success. If the\
    \ target's saving throw is successful, or if the effect ends on it, the target\
    \ is immune to the Chilling Gaze of all yetis (but not abominable yetis) for 1\
    \ hour."
  "name": "Chilling Gaze"
"source":
- "MM"
- "SKT"
- "TftYP"
- "EGW"
- "IDRotF"
- "LoX"
- "ToFW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Yeti.webp"
```
^statblock

```encounter-table
name: Yeti
creatures:
 - 1: Yeti
```

## Environment

arctic