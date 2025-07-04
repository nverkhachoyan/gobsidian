---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/21
- monster/size/large
- monster/type/fiend/devil
aliases: ["Moloch"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 177
---
# [Moloch](3-Mechanics\CLI\bestiary\npc/moloch-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 177*  

> [!quote]- A quote from Mordenkainen  
> 
> Moloch obsesses over power he lost rather than thinking of the power he could gain elsewhere in the planes. What a pity he so wastes his potential.

## Moloch

Exiled from the Nine Hells, Moloch would do anything to reclaim his position. Long ago, Moloch earned his place among the other archdevils through the glory he won driving demons out of the Nine Hells. Asmodeus rewarded him by elevating Moloch to the rulership of Malbolge.

For eons, Moloch ruled his domain, vying against the other archdevils as he sought still greater power. This animosity worked in Asmodeus's favor, since Asmodeus knew that Moloch's scheming helped keep the other archdevils in check. The arrangement began to unravel, however, when Moloch took the night hag named Malagard for his advisor. Her words were poison, and gradually she convinced Moloch to direct his efforts to topple Asmodeus.

Although the conspiracy nearly succeeded, it was thwarted. Moloch was stripped of his station and sentenced to death—and only the timely use of a planar portal allowed him to escape.

Moloch wasted no time in preparing for his return. He amassed an army of devils and monsters and left them to make final preparations for invading the Nine Hells, while he ventured to a distant Material Plane in the hope of finding an artifact that would ensure his success. While there, he became trapped, leaving his armies at the mercy of his enemies. In short order they were destroyed.

Now, Moloch has been rendered nearly powerless after his last failure. He endlessly schemes of ways to return to his former status, but every time he enters the Nine Hells, he is demoted to an imp and can't regain his normal powers until he leaves. Thus, he lives a split existence, sometimes scheming in Malbolge or other layers of the Hells and at other times wandering the planes searching for magical might or secrets that might help him win back his title.

Rumors suggest that he can often be found in Sigil, where he bargains with yugoloths to build yet another army with which he might invade Malbolge and wrest the throne from Glasya. Bereft as he is, he has little to offer in exchange, so he might bargain with mortals to gain their aid in acquiring coin, jewels, and other riches in return for knowledge about the Nine Hells and the other planes.

Most of Moloch's cultists have switched allegiance to one of the other archdevils, but idols constructed to honor him still stand in deep dungeons, their jeweled eyes and the remnants of power they hold drawing monstrous worshipers and unwise adventurers into places where his foul influence remains.

```statblock
"name": "Moloch (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "253"
"hit_dice": "22d10 + 132"
"stats":
- !!int "26"
- !!int "19"
- !!int "22"
- !!int "21"
- !!int "18"
- !!int "23"
"speed": "30 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "11"
  "Wisdom": !!int "11"
  "Constitution": !!int "13"
"skillsaves":
  "Intimidation": !!int "13"
  "Deception": !!int "13"
  "Perception": !!int "11"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 21"
"languages": "all, telepathy 120 ft."
"cr": "21"
"traits":
- "desc": "Moloch's innate spellcasting ability is Charisma (spell save DC 21). He\
    \ can innately cast the following spells, requiring no material components:\n\n\
    At will: [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium\
    \ when changing his appearance), [animate dead](/3-Mechanics/CLI/spells/animate-dead.md),\
    \ [burning hands](/3-Mechanics/CLI/spells/burning-hands.md) (as a 7th-level spell),\
    \ [confusion](/3-Mechanics/CLI/spells/confusion.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [fly](/3-Mechanics/CLI/spells/fly.md), [geas](/3-Mechanics/CLI/spells/geas.md),\
    \ [major image](/3-Mechanics/CLI/spells/major-image.md), [stinking cloud](/3-Mechanics/CLI/spells/stinking-cloud.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md), [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\
    \n1/day each: [flame strike](/3-Mechanics/CLI/spells/flame-strike.md), [symbol](/3-Mechanics/CLI/spells/symbol.md)\
    \ (stunning only)"
  "name": "Innate Spellcasting"
- "desc": "If Moloch fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Moloch has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Moloch's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Moloch regains 20 hit points at the start of his turn. If he takes radiant\
    \ damage, this trait doesn't function at the start of his next turn. Moloch dies\
    \ only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Moloch makes three attacks: one with his bite, one with his claw, and one\
    \ with his whip."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 5 ft., one\
    \ target. Hit: dice:4d8 + 8|text(26) (4d8 + 8) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 30 ft., one\
    \ target. Hit: dice:2d4 + 8|text(13) (2d4 + 8) slashing damage plus dice:2d10|text(11)\
    \ (2d10) lightning damage. If the target is a creature, it must succeed on a\
    \ DC 24 Strength saving throw or be pulled up to 30 feet in a straight line toward\
    \ Moloch."
  "name": "Many-Tailed Whip"
- "desc": "Moloch exhales in a 30-foot cube. Each creature in that area must succeed\
    \ on a DC 21 Wisdom saving throw or take dice:5d10|text(27) (5d10) psychic\
    \ damage, drop whatever it is holding, and become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, a creature must take the Dash action and move away from Moloch\
    \ by the safest available route on each of its turns, unless there is nowhere\
    \ to move, in which case it needn't take the Dash action. If the creature ends\
    \ its turn in a location where it doesn't have line of sight to Moloch, the creature\
    \ can repeat the saving throw. On a success, the effect ends."
  "name": "Breath of Despair (Recharge 5-6)"
- "desc": "Moloch magically teleports, along with any equipment he is wearing and\
    \ carrying, up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Moloch casts [stinking cloud](/3-Mechanics/CLI/spells/stinking-cloud.md)."
  "name": "Stinking Cloud"
- "desc": "Moloch uses his Teleport action."
  "name": "Teleport"
- "desc": "Moloch makes one attack with his whip."
  "name": "Whip"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Moloch.webp"
```
^statblock

```encounter-table
name: Moloch
creatures:
 - 1: Moloch
```