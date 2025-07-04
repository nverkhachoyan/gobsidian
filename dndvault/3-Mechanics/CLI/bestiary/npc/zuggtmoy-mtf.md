---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/23
- monster/size/large
- monster/type/fiend/demon
aliases: ["Zuggtmoy"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 157
---
# [Zuggtmoy](3-Mechanics\CLI\bestiary\npc/zuggtmoy-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 157*  

## Zuggtmoy

The Demon Queen of Fungi, Lady of Rot and Decay, Zuggtmoy is an alien creature whose only desire is to infect the living with spores, transforming them into her mindless servants and, eventually, into decomposing hosts for the mushrooms, molds, and other fungi that she spawns.

Utterly inhuman, Zuggtmoy can mold her fungoid form into an approximation of a humanoid shape, including the skeletal-thin figure depicted in grimoires and ancient art, draped and veiled in mycelia and lichen. Indeed, much of her appearance and manner, and that of her servants, is a soulless mockery of mortal life and its many facets.

Zuggtmoy's cultists often follow her unwittingly. Most are fungi—infected to some degree, whether through inhaling her mind-controlling spores or being transformed to the point where flesh and fungus become one. Such cultists are fungal extensions of the Demon Queen's will.

Their devotion might begin with the seemingly harmless promises offered by exotic spores and mushrooms, but quickly consumes them, body and soul.

## Zuggtmoy's Lair

Zuggtmoy's principal lair is her palace on Shedaklah. It consists of two dozen mushrooms of pale yellow and rancid brown. These massive fungi are some of the largest in existence. They are surrounded by a field of acidic puffballs and poisonous vapors. The mushrooms themselves are all interconnected by bridges of shelf-fungi, and countless chambers have been hollowed out from within their rubbery, fibrous stalks.

```statblock
"name": "Zuggtmoy (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "304"
"hit_dice": "32d10 + 128"
"stats":
- !!int "22"
- !!int "15"
- !!int "18"
- !!int "20"
- !!int "19"
- !!int "24"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "11"
  "Constitution": !!int "11"
"skillsaves":
  "Perception": !!int "11"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 21"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Zuggtmoy's spellcasting ability is Charisma (spell save DC 22). She can\
    \ innately cast the following spells, requiring no material components:\n\nAt\
    \ will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [locate animals\
    \ or plants](/3-Mechanics/CLI/spells/locate-animals-or-plants.md), [ray of sickness](/3-Mechanics/CLI/spells/ray-of-sickness.md)\n\
    \n1/day each: [etherealness](/3-Mechanics/CLI/spells/etherealness.md), [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\
    \n3/day each: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [ensnaring\
    \ strike](/3-Mechanics/CLI/spells/ensnaring-strike.md), [entangle](/3-Mechanics/CLI/spells/entangle.md),\
    \ [plant growth](/3-Mechanics/CLI/spells/plant-growth.md)"
  "name": "Innate Spellcasting"
- "desc": "If Zuggtmoy fails a saving throw, she can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Zuggtmoy has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Zuggtmoy's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "Zuggtmoy makes three pseudopod attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 6|text(15) (2d8 + 6) bludgeoning damage plus dice:2d8|text(9)\
    \ (2d8) poison damage."
  "name": "Pseudopod"
- "desc": "Zuggtmoy releases spores that burst out in a cloud that fills a 20-foot-radius\
    \ sphere centered on her, and it lingers for 1 minute. Any flesh-and-blood creature\
    \ in the cloud when it appears, or that enters it later, must make a DC 19 Constitution\
    \ saving throw. On a successful save, the creature can't be infected by these\
    \ spores for 24 hours. On a failed save, the creature is infected with a disease\
    \ called the spores of Zuggtmoy and also gains a random form of madness (determined\
    \ by rolling on the Madness of Zuggtmoy table) that lasts until the creature is\
    \ cured of the disease or dies. While infected in this way, the creature can't\
    \ be reinfected, and it must repeat the saving throw at the end of every 24 hours,\
    \ ending the infection on a success. On a failure, the infected creature's body\
    \ is slowly taken over by fungal growth, and after three such failed saves, the\
    \ creature dies and is reanimated as a spore servant if it's a type of creature\
    \ that can be (see the \"Myconids\" entry in the Monster Manual)."
  "name": "Infestation Spores (3/Day)"
- "desc": "Zuggtmoy releases spores that burst out in a cloud that fills a 20-foot-radius\
    \ sphere centered on her, and it lingers for 1 minute. Humanoids and beasts in\
    \ the cloud when it appears, or that enter it later, must make a DC 19 Wisdom\
    \ saving throw. On a successful save, the creature can't be infected by these\
    \ spores for 24 hours. On a failed save, the creature is infected with a disease\
    \ called the influence of Zuggtmoy for 24 hours. While infected in this way, the\
    \ creature is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by her and\
    \ can't be reinfected by these spores."
  "name": "Mind Control Spores (Recharge 5-6)"
"reactions":
- "desc": "When Zuggtmoy is hit by an attack, one creature within 5 feet of Zuggtmoy\
    \ that is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by her must\
    \ use its reaction to be hit by the attack instead."
  "name": "Protective Thrall"
"legendary_actions":
- "desc": "Zuggtmoy makes one pseudopod attack."
  "name": "Attack"
- "desc": "One creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by\
    \ Zuggtmoy that she can see must use its reaction to move up to its speed as she\
    \ directs or to make a weapon attack against a target that she designates."
  "name": "Exert Will"
"lair_actions":
- "desc": "On Initiative count 20 (losing initiative ties), Zuggtmoy can take a lair\
    \ action to cause one of the following effects; she can't use the same effect\
    \ two rounds in a row:"
  "name": ""
- "desc": "- Zuggtmoy causes four [gas spores](/3-Mechanics/CLI/bestiary/plant/gas-spore.md)\
    \ or [violet fungi](/3-Mechanics/CLI/bestiary/plant/violet-fungus.md) to appear\
    \ in unoccupied spaces that she chooses within the lair. They vanish after 1 hour.\
    \  \n- Up to four plant creatures that are friendly to Zuggtmoy and that Zuggtmoy\
    \ can see can use their reactions to move up to their speed and make one weapon\
    \ attack.  \n- Zuggtmoy uses either her Infestation Spores or her Mind Control\
    \ Spores, centered on a mushroom or other fungus within her lair, instead of on\
    \ herself.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Zuggtmoy's lair is warped by his magic, creating\
    \ one or more of the following effects:"
  "name": ""
- "desc": "- Molds and fungi grow on surfaces within 6 miles of the lair, even where\
    \ they would normally find no purchase.  \n- Plant life within 1 mile of the lair\
    \ becomes infested with parasitic fungi, slowly mutating as it is overwhelmed.\
    \  \n- If a humanoid spends at least 1 hour within 1 mile of the lair, that creature\
    \ must succeed on a DC 17 Wisdom saving throw or descend into a madness determined\
    \ by the Madness of Zuggtmoy table. A creature that succeeds on this saving throw\
    \ can't be affected by this regional effect again for 24 hours.  "
  "name": ""
- "desc": "If Zuggtmoy dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Zuggtmoy's lair or within line of sight of the\
    \ demon lord, roll on the Madness of Zuggtmoy table to determine the nature of\
    \ the madness, which is a character flaw that lasts until cured. See the \"Dungeon\
    \ Master's Guide\" for more on madness.\n\nMadness of Zuggtmoy\n\ndice: [](zuggtmoy-mtf.md#^madness-of-zuggtmoy)\n\
    \n| dice: d100 | Flaw (lasts until cured) |\n|------------|--------------------------|\n\
    | 01–20 | \"I see visions in the world around me that others do not.\" |\n| 21–\
    40 | \"I periodically slip into a catatonic state, staring off into the distance\
    \ for long stretches at a time.\" |\n| 41–60 | \"I see an altered version of reality,\
    \ with my mind convincing itself that things are true even in the face of overwhelming\
    \ evidence to the contrary.\" |\n| 61–80 | \"My mind is slipping away, and my\
    \ intelligence seems to wax and wane.\" |\n| 81–00 | \"I am constantly scratching\
    \ at unseen fungal infections.\" |\n^madness-of-zuggtmoy"
  "name": "Madness of Zuggtmoy"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Zuggtmoy.webp"
```
^statblock

```encounter-table
name: Zuggtmoy
creatures:
 - 1: Zuggtmoy
```