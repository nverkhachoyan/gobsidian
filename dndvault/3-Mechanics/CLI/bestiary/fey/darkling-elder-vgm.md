---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fey
aliases: ["Darkling Elder"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 134
---
# [Darkling Elder](3-Mechanics\CLI\bestiary\fey/darkling-elder-vgm.md)
*Source: Volo's Guide to Monsters p. 134*  

Ancient legends speak of a seelie fey who betrayed the Summer Queen. His true name has been stricken from history, but the stories call him Dubh Catha ("Dark Crow" in Common). So great was the Summer Queen's wrath that she cursed every member of his house. Other fey refer to the descendants of Dubh Catha's house as the dubh sith-or, in Common, "darklings." Darklings most often settle in secluded caverns and chambers beneath the towns of other species. From such enclaves, they quietly ply their trade as thieves and assassins.

## The Killing Light

The Summer Queen's curse causes a darkling's body to absorb light, and doing so wizens the creature, much like the effect of rapid aging. For this reason, darklings cover every part of their body with clothing when exposure to light is a risk. The light a darkling absorbs over the course of its lifetime explodes outward when the darkling dies, incinerating the creature and much of its possessions.

## Love of Art

Despite their curse, darklings retain a fondness for the beauty of art. A darkling might risk taking a peek at a sunset or lighting a tiny candle to glimpse the colors in a painting or a jewel.

## Elder Transformation

A wise and respected darkling can qualify to undergo a ritual to become an elder. Other elders mark the supplicant with glowing tattoos, channeling some of the darkling's absorbed light away from its body. If the ritual succeeds, the darkling grows into a tall and fair form, like that of a gray-skinned elf. The darkling perishes if the ritual fails.

```statblock
"name": "Darkling Elder (VGM)"
"size": "Medium"
"type": "fey"
"alignment": "Chaotic Neutral"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "27"
"hit_dice": "5d8 + 5"
"stats":
- !!int "13"
- !!int "17"
- !!int "12"
- !!int "10"
- !!int "14"
- !!int "13"
"speed": "30 ft."
"skillsaves":
  "Deception": !!int "3"
  "Stealth": !!int "7"
  "Perception": !!int "6"
  "Acrobatics": !!int "5"
"senses": "blindsight 30 ft., darkvision 120 ft., passive Perception 16"
"languages": "Elvish, Sylvan"
"cr": "2"
"traits":
- "desc": "When the darkling elder dies, magical light flashes out from it in a 10-foot\
    \ radius as its body and possessions, other than metal or magic objects, burn\
    \ to ash. Any creature in that area must make a DC 11 Constitution saving throw.\
    \ On a failure, the creature takes dice:2d6|text(7) (2d6) radiant damage and,\
    \ if the creature can see the light, is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ until the end of its next turn. If the saving throw is successful, the creature\
    \ takes half the damage and isn't [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)."
  "name": "Death Burn"
"actions":
- "desc": "The darkling elder makes two melee attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage. If the darkling elder\
    \ had advantage on the attack roll, the attack deals as: extra dice:3d6|text(10)\
    \ (3d6) piercing damage."
  "name": "Shortsword"
- "desc": "The darkling elder casts [darkness](/3-Mechanics/CLI/spells/darkness.md)\
    \ without any components. Wisdom is its spellcasting ability."
  "name": "Darkness (Recharges after a Short or Long Rest)"
"source":
- "VGM"
- "WBtW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Darkling%20Elder.webp"
```
^statblock

```encounter-table
name: Darkling Elder
creatures:
 - 1: Darkling Elder
```

## Environment

underdark, forest, swamp, urban