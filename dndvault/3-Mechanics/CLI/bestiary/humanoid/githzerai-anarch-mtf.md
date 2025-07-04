---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/16
- monster/size/medium
- monster/type/humanoid/gith
aliases: ["Githzerai Anarch"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 207
---
# [Githzerai Anarch](3-Mechanics\CLI\bestiary\humanoid/githzerai-anarch-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 207*  

> [!quote]- A quote from Mordenkainen  
> 
> The githzerai are a check on the githyanki and the illithids. The githyanki are a check on the githzerai and the illithids. Thus, three unequal forces enforce the Balance.

## Githzerai Anarch

The most powerful of the githzerai, anarchs lead communities and maintain the adamantine citadels that serve as strong points in planes beyond Limbo. They have formidable psionic capabilities, able to manipulate the unformed substance of their adopted plane with a thought. These rare githzerai are sages and mystics, and their word is law.

## Gith

The descendants of an ancient people—so old their original name has been lost—have turned against each other, becoming vicious enemies divided over mortality, purpose, and the machinations of their leaders. The bellicose githyanki terrorize the Astral Plane, raiding into other worlds to plunder the multiverse of its magic and riches. The githzerai live apart from the rest of the cosmos, content within the confines of their fortresses floating through the chaos of Limbo. Although the two groups of gith despise each other, their hatred for the mind flayers from whom they escaped endures, and both githyanki and githzerai are dedicated to hunting their ancestral foes.

> [!quote]- A quote from Mordenkainen  
> 
> What would become of this multiverse if githyanki didn't guard the Astral Plane from the illithid menace? What would reality become if beings of thought ruled the plane of thought?

## An Anarch's Lair

In Limbo, githzerai anarchs create islands of tranquility in the otherwise turbulent plane. By directing its psionic power, an anarch can give form to formless substance, creating mountains, lakes, and structures of any composition to serve as a foundation for a githzerai community.

```statblock
"name": "Githzerai Anarch (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Lawful Neutral"
"ac": !!int "20"
"hp": !!int "144"
"hit_dice": "17d8 + 68"
"stats":
- !!int "16"
- !!int "21"
- !!int "18"
- !!int "18"
- !!int "20"
- !!int "14"
"speed": "30 ft., fly 40 ft. (hover)"
"saves":
  "Dexterity": !!int "10"
  "Wisdom": !!int "10"
  "Intelligence": !!int "9"
  "Strength": !!int "8"
"skillsaves":
  "Insight": !!int "10"
  "Perception": !!int "10"
  "Arcana": !!int "9"
"senses": "passive Perception 20"
"languages": "Gith"
"cr": "16"
"traits":
- "desc": "The anarch's innate spellcasting ability is Wisdom (spell save DC 18, dice:\
    \ d20+10 (+10) to hit with spell attacks). It can innately cast the following\
    \ spells, requiring no components:\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\
    \ (the hand is invisible)\n\n1/day each: [globe of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md),\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md), [teleportation circle](/3-Mechanics/CLI/spells/teleportation-circle.md),\
    \ [wall of force](/3-Mechanics/CLI/spells/wall-of-force.md)\n\n3/day each:\
    \ [feather fall](/3-Mechanics/CLI/spells/feather-fall.md), [jump](/3-Mechanics/CLI/spells/jump.md),\
    \ [see invisibility](/3-Mechanics/CLI/spells/see-invisibility.md), [shield](/3-Mechanics/CLI/spells/shield.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "While the anarch is wearing no armor and wielding no shield, its AC includes\
    \ its Wisdom modifier."
  "name": "Psychic Defense"
"actions":
- "desc": "The anarch makes three unarmed strikes."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d8 + 5|text(14) (2d8 + 5) bludgeoning damage plus dice:4d8|text(18)\
    \ (4d8) psychic damage."
  "name": "Unarmed Strike"
"legendary_actions":
- "desc": "The anarch makes one unarmed strike."
  "name": "Strike"
- "desc": "The anarch magically teleports, along with any equipment it is wearing\
    \ and carrying, to an unoccupied space it can see within 30 feet of it."
  "name": "Teleport"
- "desc": "The anarch casts the [reverse gravity](/3-Mechanics/CLI/spells/reverse-gravity.md)\
    \ spell. The spell has the normal effect, except that the anarch can orient the\
    \ area in any direction and creatures and objects fall toward the end of the area."
  "name": "Change Gravity (Costs 3 Actions)"
"lair_actions":
- "desc": "An anarch can use lair actions. On initiative count 20 (losing initiative\
    \ ties), the anarch can take a lair action to cause one of the following effects;\
    \ the anarch can't use the same effect two rounds in a row:"
  "name": ""
- "desc": "- The anarch casts the [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)\
    \ spell (at 5th level), but the anarch can change the damage type from lightning\
    \ to cold, fire, psychic, radiant, or thunder. If the spell deals damage other\
    \ than fire or lightning, it doesn't ignite flammable objects.  \n- The anarch\
    \ casts the [creation](/3-Mechanics/CLI/spells/creation.md) spell (as a 9th-level\
    \ spell) using the unformed substance of Limbo instead of shadow material. If\
    \ used in Limbo, the object remains until the anarch's [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ is broken, regardless of its composition. If the anarch moves more than 120\
    \ feet from the object, its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ breaks.  \n- The anarch can magically move an object it can see within 150 feet\
    \ of it by making a Wisdom check with advantage. The DC depends on the object's\
    \ size: DC 5 for Tiny, DC 10 for Small, DC 15 for Medium, DC 20 for Large, and\
    \ DC 25 for Huge or larger.  "
  "name": ""
"regional_effects":
- "desc": "The region containing an anarch's lair is warped by its presence, which\
    \ creates one or more of the following effects:"
  "name": ""
- "desc": "- In Limbo, the anarch can spend 10 minutes stabilizing a 5-mile area centered\
    \ on it, causing the unformed substance to take whatever inanimate form the anarch\
    \ chooses. During that process, the anarch determines the shape and composition\
    \ of the forms created.  \n- The anarch stabilizes any object created in Limbo\
    \ and brought to the Material Plane for as long as the anarch remains within 1\
    \ mile of it (no action required).  "
  "name": ""
- "desc": "If the anarch dies, these effects end after dice: 1d6|avg|noform (1d6)\
    \ rounds. All formed substance becomes a chaotic churn of energy and matter, unraveling\
    \ into unformed substance that dissipates dice: 1d6|avg|noform (1d6) rounds\
    \ later."
  "name": ""
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Githzerai%20Anarch.webp"
```
^statblock

```encounter-table
name: Githzerai Anarch
creatures:
 - 1: Githzerai Anarch
```