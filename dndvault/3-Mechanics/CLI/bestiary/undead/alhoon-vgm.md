---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/10
- monster/environment/underdark
- monster/size/medium
- monster/type/undead
aliases: ["Alhoon"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 172
---
# [Alhoon](3-Mechanics\CLI\bestiary\undead/alhoon-vgm.md)
*Source: Volo's Guide to Monsters p. 172*  

Mind flayers that pursue arcane magic are exiled as deviants, and for them no eternal communion with an elder brain is possible. The road to lichdom offers a way to escape the permanency of death, but that path is long and solitary. Alhoons are mind flayers that use a shortcut.

## Arcane Temptation

Elder brains forbid mind flayers from pursuing magic power aside from psionics, but it isn't an interdiction they must often enforce. Illithids brook no masters but members of their own kind, so it isn't in their nature to bow to any god or otherworldly patron. However, wizardry remains a rare temptation.

In the pages of a spellbook, an illithid sees a system to acquire authority. Through the writings of the wizard who penned it, the illithid perceives the workings of a highly intelligent mind. Most mind flayers who find a spellbook react with abhorrence or indifference, but for some a spellbook is a gateway to a new way of thinking.

For a time, the study of such forbidden texts can be hidden from other illithids and even from an elder brain. Understanding of wizardry eludes the mind like a living thing. Yet eventually, understanding comes, and a mind flayer arcanist must accept itself as deviant and flee the colony if it is to live.

Confronting this awful reality, a group of nine mind flayer deviants used their arcane magic and psionics to weave a new truth. These nine called themselves the alhoon, and ever afterward, all those who follow in their footsteps have been referred to by the same name.

Initially, an alhoon can be difficult to distinguish from a normal mind flayer. The most obvious difference is the lack of the mind flayer's ever-present mucus coating. Without that protection, an alhoon's skin becomes dry and cracked. Its eyes might appear shriveled and sunken. Both of these clues are easily missed by someone who hasn't seen a mind flayer. However, in short order, an alhoon's flesh withers away and its empty eye sockets gleam with cold pinpricks of light like other liches.

The undeath conferred by a periapt of mind trapping lasts only so long as the life of the living victim selected. Thus an alhoon who brought a 200-year-old elf to be sacrificed looks forward to a much longer existence than one that sacrifices a 35-year-old person. Alhoons can extend their existence by repeating the ritual with new victims, effectively resetting the clocks for themselves.

Destruction of a periapt of mind trapping consigns those trapped within it to oblivion, and thus alhoons often work together to create elaborate protections about the periapt and their preferred ritual site. Sometimes a single alhoon is entrusted with the periapt of mind trapping, but this is a dangerous proposition. Anyone who holds the periapt of mind trapping gains advantage on attacks, saves, and check against the alhoons associated with its creation, and those alhoons in turn suffer disadvantage on attacks, saves, and check against the holder.

In addition, the holder of the periapt can telepathically communicate with any sacrificed soul trapped within, and alhoons within the periapt can speak telepathically with the holder. A creature carrying the periapt can't prevent communication from alhoons but can silence trapped souls.

## Existential Fear

Arcanist deviants that taste freedom from the colony react in a variety of ways. Some prize their privacy, others seek to commune with similar minds, and still others seek to dominate a colony, elevating themselves to the position of leadership normally held by an elder brain. Regardless of the arcanist's personal inclinations, it faces the same stark fact: When it dies, it will not join the host of minds in the elder brain. Deviant minds are never accepted as part of the collective. For it, death means oblivion.

## Dreadful Deliverance

Lichdom offers salvation and the prospect of being able to pursue knowledge indefinitely. Having feasted on the brains of people when alive, a mind flayer has no compunction about feeding souls to a phylactery. The only hindrance to a mind flayer becoming a lich is the means, which is a secret some mind flayer arcanists stop at nothing to discover. Yet lichdom requires an arcane spellcaster to be at the apex of power, something many mind flayers find is far from their grasps.

## A Psionic Secret

Alhoons can cooperate in the creation of a periapt of mind trapping, a fist-sized container made of silver, emerald, and amethyst. The process requires at least three mind flayer arcanists and the sacrifice of an equal number of souls from living victims in a three-day-long ritual of spellcasting and psionic communion. Upon its completion, free-willed undeath is conferred on the mind flayers, turning them into alhoons.

## Precarious Immortality

Unlike with true lichdom, the periapt of mind trapping doesn't restore the alhoons to undeath if they are destroyed. Instead, a destroyed alhoon's mind is transferred to the periapt where it remains in communion with any other trapped alhoon minds, as well as the souls of those sacrificed.

```statblock
"name": "Alhoon (VGM)"
"size": "Medium"
"type": "undead"
"alignment": "Any Evil alignment"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "120"
"hit_dice": "16d8 + 48"
"stats":
- !!int "11"
- !!int "12"
- !!int "16"
- !!int "19"
- !!int "17"
- !!int "17"
"speed": "30 ft."
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "7"
  "Intelligence": !!int "8"
  "Constitution": !!int "7"
"skillsaves":
  "Deception": !!int "7"
  "Stealth": !!int "5"
  "Insight": !!int "7"
  "Perception": !!int "7"
  "History": !!int "8"
  "Arcana": !!int "8"
"damage_resistances": "cold, lightning, necrotic"
"damage_immunities": "poison; bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 17"
"languages": "Deep Speech, Undercommon, telepathy 120 ft."
"cr": "10"
"traits":
- "desc": "The alhoon's innate spellcasting ability is Intelligence (spell save DC\
    \ 16). It can innately cast the following spells, requiring no components:\n\n\
    At will: [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [levitate](/3-Mechanics/CLI/spells/levitate.md)\n\
    \n1/day each: [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md),\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "The alhoon is a 12th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 16, dice: d20+8 (+8) to hit with spell attacks). The alhoon\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [chill\
    \ touch](/3-Mechanics/CLI/spells/chill-touch.md), [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md),\
    \ [shocking grasp](/3-Mechanics/CLI/spells/shocking-grasp.md)\n\n1st level (4\
    \ slots): [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [disguise\
    \ self](/3-Mechanics/CLI/spells/disguise-self.md), [magic missile](/3-Mechanics/CLI/spells/magic-missile.md),\
    \ [shield](/3-Mechanics/CLI/spells/shield.md)\n\n2nd level (3 slots): [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [mirror image](/3-Mechanics/CLI/spells/mirror-image.md), [scorching ray](/3-Mechanics/CLI/spells/scorching-ray.md)\n\
    \n3rd level (3 slots): [counterspell](/3-Mechanics/CLI/spells/counterspell.md),\
    \ [fly](/3-Mechanics/CLI/spells/fly.md), [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)\n\
    \n4th level (3 slots): [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [Evard's black tentacles](/3-Mechanics/CLI/spells/evards-black-tentacles.md),\
    \ [phantasmal killer](/3-Mechanics/CLI/spells/phantasmal-killer.md)\n\n5th level\
    \ (2 slots): [modify memory](/3-Mechanics/CLI/spells/modify-memory.md), [wall\
    \ of force](/3-Mechanics/CLI/spells/wall-of-force.md)\n\n6th level (1 slots):\
    \ [disintegrate](/3-Mechanics/CLI/spells/disintegrate.md), [globe of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md)"
  "name": "Spellcasting"
- "desc": "The alhoon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The alhoon has advantage on saving throws against any effect that turns\
    \ undead."
  "name": "Turn Resistance"
"actions":
- "desc": "Melee Spell Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6|text(10) (3d6) cold damage."
  "name": "Chilling Grasp"
- "desc": "The alhoon magically emits psychic energy in a 60-foot cone. Each creature\
    \ in that area must succeed on a DC 16 Intelligence saving throw or take dice:4d8\
    \ + 4|text(22) (4d8 + 4) psychic damage and be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. A target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Mind Blast (Recharge 5-6)"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Alhoon.webp"
```
^statblock

```encounter-table
name: Alhoon
creatures:
 - 1: Alhoon
```

## Environment

underdark