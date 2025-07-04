---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/21
- monster/size/large
- monster/type/fiend/devil
aliases: ["Hutijin"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 175
---
# [Hutijin](3-Mechanics\CLI\bestiary\npc/hutijin-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 175*  

> [!quote]- A quote from Mordenkainen  
> 
> What price would you put on your life? How much then, for your soul? Bargain with Hutijin, and it will cost you both. I hope it is worth it.

## Hutijin

Politics in the Nine Hells are anything but predictable. Alliances form all the time, but most wind up unraveling due to treachery. Nevertheless, for all their backbiting and betrayal, the devils do occasionally display loyalty, offering unwavering service to their masters. One such example is Hutijin, a duke of Cania and loyal servant of Mephistopheles.

Across the Hells, Hutijin's name fills lesser devils with fear and loathing, for this duke commands two companies of pit fiends, which make up Cania's aristocracy. With such soldiers under his command, Hutijin can easily crush any rival who gets in his way, while also providing Mephistopheles with security against armies that might seek to contest his dominion. Hutijin has amassed enough power to challenge the lord of Cania, but he has never wavered in his support for his master—suggesting, perhaps, that Mephistopheles has some hold over him.

Outside the Nine Hells, Hutijin is a relatively obscure figure, known only to the most learned infernal scholars. He has no cults of his own, and his servants are few in number. The reason is simple: Hutijin hates mortals. When summoned from the Hells, he repays the instigator with a long and agonizing death.

Mephistopheles forbids Hutijin from making too many forays into the Material Plane, since the duke's absence leaves him vulnerable to his rivals. Other archdevils know how much Hutijin despises mortals and have secretly disseminated the means to call him from the Nine Hells in the hope of distracting the archdevil long enough for them to assail Mephistopheles. Hutijin sends devils into the Material Plane to eradicate mention of his name and destroy those who have learned of him, but the summonings still occur. When called from his post, he negotiates as quickly as he can, usually closing a deal with little cost to the summoner. However, once the deal has been struck, Hutijin repays the interruption with death.

```statblock
"name": "Hutijin (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "200"
"hit_dice": "16d10 + 112"
"stats":
- !!int "27"
- !!int "15"
- !!int "25"
- !!int "23"
- !!int "19"
- !!int "25"
"speed": "30 ft., fly 60 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "11"
  "Constitution": !!int "14"
"skillsaves":
  "Intimidation": !!int "14"
  "Perception": !!int "11"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 21"
"languages": "all, telepathy 120 ft."
"cr": "21"
"traits":
- "desc": "Hutijin's innate spellcasting ability is Charisma (spell save DC 22). He\
    \ can innately cast the following spells, requiring no material components:\n\n\
    At will: [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium\
    \ when changing his appearance), [animate dead](/3-Mechanics/CLI/spells/animate-dead.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [hold monster](/3-Mechanics/CLI/spells/hold-monster.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md) (self only), [lightning\
    \ bolt](/3-Mechanics/CLI/spells/lightning-bolt.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md),\
    \ [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\n1/day each: [heal](/3-Mechanics/CLI/spells/heal.md),\
    \ [symbol](/3-Mechanics/CLI/spells/symbol.md) (hopelessness only)\n\n3/day:\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md)"
  "name": "Innate Spellcasting"
- "desc": "Each creature within 15 feet of Hutijin that isn't a devil makes saving\
    \ throws with disadvantage."
  "name": "Infernal Despair"
- "desc": "If Hutijin fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Hutijin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Hutijin's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Hutijin regains 20 hit points at the start of his turn. If he takes radiant\
    \ damage, this trait doesn't function at the start of his next turn. Hutijin dies\
    \ only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Hutijin makes four attacks: one with his bite, one with his claw, one with\
    \ his mace, and one with his tail."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 8|text(15) (2d6 + 8) piercing damage. The target\
    \ must succeed on a DC 22 Constitution saving throw or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\
    \ While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way,\
    \ the target can't regain hit points, and it takes dice:3d6|text(10) (3d6)\
    \ poison damage at the start of each of its turns. The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 8|text(15) (2d6 + 8) bludgeoning damage."
  "name": "Mace"
- "desc": "Melee Weapon Attack: dice: d20+15 (+15) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d10 + 8|text(19) (2d10 + 8) bludgeoning damage."
  "name": "Tail"
- "desc": "Hutijin magically teleports, along with any equipment he is wearing and\
    \ carrying, up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"reactions":
- "desc": "In response to taking damage, Hutijin utters a dreadful word of power.\
    \ Each creature within 30 feet of him that isn't a devil must succeed on a DC\
    \ 22 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of him for 1 minute. A creature can repeat the saving throw at the end of each\
    \ of its turns, ending the effect on itself on a success. A creature that saves\
    \ against this effect is immune to Hutijin's Fearful Voice for 24 hours."
  "name": "Fearful Voice (Recharge 5-6)"
"legendary_actions":
- "desc": "Hutijin attacks once with his mace."
  "name": "Attack"
- "desc": "Hutijin releases lightning in a 20-foot radius. All other creatures in\
    \ that area must each make a DC 22 Dexterity saving throw, taking dice:4d8|text(18)\
    \ (4d8) lightning damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Lightning Storm (Costs 2 Actions)"
- "desc": "Hutijin uses his Teleport action."
  "name": "Teleport"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Hutijin.webp"
```
^statblock

```encounter-table
name: Hutijin
creatures:
 - 1: Hutijin
```