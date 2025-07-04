---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/20
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Drow Matron Mother"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 186
---
# [Drow Matron Mother](3-Mechanics\CLI\bestiary\humanoid/drow-matron-mother-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 186*  

> [!quote]- A quote from Mordenkainen  
> 
> Matron mother is a strange title for a cruel tyrant, but given what drow consider to be a goddess, perhaps we shouldn't be surprised.

## Drow Matron Mother

At the head of each drow noble house sits a matron mother, an influential priestess of Lolth charged with carrying out the god's will while also advancing the interests of the family. Matron mothers embody the scheming and treachery associated with the Queen of Spiders. Each stands at the center of a vast conspiratorial web, with demons, drow, spiders, and slaves positioned between them and their enemies. Although matron mothers command great power, that power depends on maintaining the Spider Queen's favor, and the dark god sometimes capriciously takes back what she has given. The stat block here represents a matron mother at the height of her power.

```statblock
"name": "Drow Matron Mother (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral Evil"
"ac": !!int "17"
"ac_class": "[half plate armor](/3-Mechanics/CLI/items/half-plate-armor.md)"
"hp": !!int "262"
"hit_dice": "35d8 + 105"
"stats":
- !!int "12"
- !!int "18"
- !!int "16"
- !!int "17"
- !!int "21"
- !!int "22"
"speed": "30 ft."
"saves":
  "Charisma": !!int "12"
  "Wisdom": !!int "11"
  "Constitution": !!int "9"
"skillsaves":
  "Stealth": !!int "10"
  "Religion": !!int "9"
  "Insight": !!int "11"
  "Perception": !!int "11"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 21"
"languages": "Elvish, Undercommon"
"cr": "20"
"traits":
- "desc": "The drow's innate spellcasting ability is Charisma (spell save DC 20).\
    \ She can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [detect\
    \ magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n1/day each: [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md),\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Innate Spellcasting"
- "desc": "The drow is a 20th-level spellcaster. Her spellcasting ability is Wisdom\
    \ (spell save DC 19, dice: d20+11 (+11) to hit with spell attacks). The drow\
    \ has the following cleric spells prepared:\n\nCantrips (at will): [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [resistance](/3-Mechanics/CLI/spells/resistance.md),\
    \ [sacred flame](/3-Mechanics/CLI/spells/sacred-flame.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1st level (4 slots): [bane](/3-Mechanics/CLI/spells/bane.md), [command](/3-Mechanics/CLI/spells/command.md),\
    \ [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md), [guiding bolt](/3-Mechanics/CLI/spells/guiding-bolt.md)\n\
    \n2nd level (3 slots): [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [silence](/3-Mechanics/CLI/spells/silence.md), [spiritual weapon](/3-Mechanics/CLI/spells/spiritual-weapon.md)\n\
    \n3rd level (3 slots): [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md),\
    \ [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [spirit guardians](/3-Mechanics/CLI/spells/spirit-guardians.md)\n\n4th level\
    \ (3 slots): [banishment](/3-Mechanics/CLI/spells/banishment.md), [death ward](/3-Mechanics/CLI/spells/death-ward.md),\
    \ [freedom of movement](/3-Mechanics/CLI/spells/freedom-of-movement.md), [guardian\
    \ of faith](/3-Mechanics/CLI/spells/guardian-of-faith.md)\n\n5th level (3 slots):\
    \ [contagion](/3-Mechanics/CLI/spells/contagion.md), [flame strike](/3-Mechanics/CLI/spells/flame-strike.md),\
    \ [geas](/3-Mechanics/CLI/spells/geas.md), [mass cure wounds](/3-Mechanics/CLI/spells/mass-cure-wounds.md)\n\
    \n6th level (2 slots): [blade barrier](/3-Mechanics/CLI/spells/blade-barrier.md),\
    \ [harm](/3-Mechanics/CLI/spells/harm.md)\n\n7th level (2 slots): [divine\
    \ word](/3-Mechanics/CLI/spells/divine-word.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\n\
    \n8th level (1 slots): [holy aura](/3-Mechanics/CLI/spells/holy-aura.md)\n\
    \n9th level (1 slots): [gate](/3-Mechanics/CLI/spells/gate.md)"
  "name": "Spellcasting"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "As a bonus action, the matron can bestow the Spider Queen's blessing on\
    \ one ally she can see within 30 feet of her. The ally takes dice:2d6|text(7)\
    \ (2d6) psychic damage but has advantage on the next attack roll it makes until\
    \ the end of its next turn."
  "name": "Lolth's Fickle Favor"
- "desc": "The drow has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The matron mother makes two demon staff attacks or three tentacle rod attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) bludgeoning damage, or dice:1d8\
    \ + 4|text(8) (1d8 + 4) bludgeoning damage if used with two hands, plus dice:4d6|text(14)\
    \ (4d6) psychic damage. In addition, the target must succeed on a DC19 Wisdom\
    \ saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the drow for 1 minute. The [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Demon Staff"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 15 ft., one target.\
    \ Hit: dice:1d6|text(3) (1d6) bludgeoning damage. If the target is hit three\
    \ times by the rod on one turn, the target must succeed on a DC 15 Constitution\
    \ saving throw or suffer the following effects for 1 minute: the target's speed\
    \ is halved, it has disadvantage on Dexterity saving throws, and it can't use\
    \ reactions. Moreover, on each of its turns, it can take either an action or a\
    \ bonus action, but not both. At the end of each of its turns, it can repeat the\
    \ saving throw, ending the effect on itself on a success."
  "name": "Tentacle Rod"
- "desc": "The drow magically summons a [retriever](/3-Mechanics/CLI/bestiary/construct/retriever-mtf.md)\
    \ or a [yochlol](/3-Mechanics/CLI/bestiary/fiend/yochlol.md). The summoned creature\
    \ appears in an unoccupied space within 60 feet of its summoner, acts as an ally\
    \ of its summoner, and can't summon other demons. It remains for 10 minutes, until\
    \ it or its summoner dies, or until its summoner dismisses it as an action."
  "name": "Summon Servant (1/Day)"
"legendary_actions":
- "desc": "The drow makes one attack with her demon staff."
  "name": "Demon Staff"
- "desc": "An allied demon within 30 feet of the drow uses its reaction to make one\
    \ attack against a target of the drow's choice that she can see."
  "name": "Compel Demon (Costs 2 Actions)"
- "desc": "The drow expends a spell slot to cast a 1st-, 2nd-, or 3rd-level spell\
    \ that she has prepared. Doing so costs 1 legendary action per level of the spell."
  "name": "Cast a Spell (Costs 1–3 Actions)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Drow%20Matron%20Mother.webp"
```
^statblock

```encounter-table
name: Drow Matron Mother
creatures:
 - 1: Drow Matron Mother
```

## Environment

underdark