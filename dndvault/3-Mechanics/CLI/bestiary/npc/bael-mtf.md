---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/19
- monster/size/large
- monster/type/fiend/devil
aliases: ["Bael"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 170
---
# [Bael](3-Mechanics\CLI\bestiary\npc/bael-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 170*  

> [!quote]- A quote from Mordenkainen  
> 
> Bael, Geryon, Hutijin, Moloch—they are pawns. Even the lords of layers such as Zariel and Mammon are merely more powerful pieces. Archdevils may make dramatic moves, but it is Asmodeus who plays the game best.

## Bael

With the Blood War raging for eons and no end in sight, opportunities abound for ambitious archdevils to win fame, glory, and power in the ongoing struggle against the demons. Duke Bael, one of Mammon's most important vassals, has won fame and acclaim for his victories. Charged with leading sixty-six companies of [barbed devils](/3-Mechanics/CLI/bestiary/fiend/barbed-devil.md), Bael has proven to be a tactical genius, earning esteem for himself and his master as a result of victory after victory over the abyssal host. Mammon relies on Bael, because of his battle acumen, to safeguard his holdings. Mammon has never been ousted during a time when so many other archdevils have lost their positions, which is a testament to Bael's skill on the battlefield.

For his accomplishments, as well as for the hue of his skin, Baal has been granted the title of Bronze General. His accolades notwithstanding, Bael has had a difficult time navigating the quagmire of infernal politics. His critics call him naive, though never to his face. His primary interest has always been leading soldiers in battle, so he finds it frustrating to have his ambitions of ascending to a higher rank constantly stymied by politically shrewd rivals.

Bael prefers to make servants out of his adversaries, and mortals bound to his service earn their wretched place by falling victim to Bael's superior stratagems. Bael gladly spares the lives of those he defeats, but only if they pledge their souls and service to him. Although he is willing to corrupt almost any being in this way, he always destroys any demons he defeats.

Bael also welcomes mortals into his service if they can provide him with an advantage in his own politicking. He recruits savvy individuals and relies on them to represent his interests at Mammon's court, which leaves Bael free to pursue his battle lust.

Despite his lack of interest in affairs outside battle, or perhaps because of it, Bael has gained a small following of cultists. Those who worship at his altar call him the King of Hell, and the most deluded believe that he is the lord of all devils. In arcane circles, certain writings, such as the dreaded Book of Fire, say that Bael revealed the [invisibility](/3-Mechanics/CLI/spells/invisibility.md) spell to the world, though some scholars of magic hotly refute such claims. Bael is sometimes depicted as a toad, a cat, a male human, or some combination of these forms, though none of these images reflect his true appearance.

```statblock
"name": "Bael (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "189"
"hit_dice": "18d10 + 90"
"stats":
- !!int "24"
- !!int "17"
- !!int "20"
- !!int "21"
- !!int "24"
- !!int "24"
"speed": "30 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "9"
  "Intelligence": !!int "11"
  "Constitution": !!int "11"
"skillsaves":
  "Intimidation": !!int "13"
  "Perception": !!int "13"
  "Persuasion": !!int "13"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 23"
"languages": "all, telepathy 120 ft."
"cr": "19"
"traits":
- "desc": "Bael's innate spellcasting ability is Charisma (spell save DC 21, dice:\
    \ d20+13 (+13) to hit with spell attacks). He can innately cast the following\
    \ spells, requiring no material components:\n\nAt will: [alter self](/3-Mechanics/CLI/spells/alter-self.md)\
    \ (can become Medium when changing his appearance), [animate dead](/3-Mechanics/CLI/spells/animate-dead.md),\
    \ [charm person](/3-Mechanics/CLI/spells/charm-person.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [inflict wounds](/3-Mechanics/CLI/spells/inflict-wounds.md) (as an 8th-level\
    \ spell), [invisibility](/3-Mechanics/CLI/spells/invisibility.md) (self only),\
    \ [major image](/3-Mechanics/CLI/spells/major-image.md)\n\n1/day each: [dominate\
    \ monster](/3-Mechanics/CLI/spells/dominate-monster.md), [symbol](/3-Mechanics/CLI/spells/symbol.md)\
    \ (stunning only)\n\n3/day each: [counterspell](/3-Mechanics/CLI/spells/counterspell.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md), [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)"
  "name": "Innate Spellcasting"
- "desc": "Bael can use a bonus action to appear dreadful until the start of his next\
    \ turn. Each creature, other than a devil, that starts its turn within 10 feet\
    \ of Bael must succeed on a DC 22 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the start of the creature's next turn."
  "name": "Dreadful"
- "desc": "If Bael fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Bael has advantage on saving throws against spells and other magical effects."
  "name": "Magic Resistance"
- "desc": "Bael's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Bael regains 20 hit points at the start of his turn. If he takes cold or\
    \ radiant damage, this trait doesn't function at the start of his next turn. Bael\
    \ dies only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Bael makes two melee attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 20 ft., one\
    \ target. Hit: dice:2d8 + 7|text(16) (2d8 + 7) piercing damage plus dice:3d8|text(13)\
    \ (3d8) necrotic damage."
  "name": "Hellish Morningstar"
- "desc": "Each ally of Bael's within 60 feet of him can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until the end\
    \ of his next turn."
  "name": "Infernal Command"
- "desc": "Bael magically teleports, along with any equipment he is wearing and carrying,\
    \ up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Bael attacks once with his hellish morningstar."
  "name": "Attack (Cost 2 Actions)"
- "desc": "Bael casts [charm person](/3-Mechanics/CLI/spells/charm-person.md) or major\
    \ image."
  "name": "Awaken Greed"
- "desc": "Bael uses his Infernal Command action."
  "name": "Infernal Command"
- "desc": "Bael uses his Teleport action."
  "name": "Teleport"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Bael.webp"
```
^statblock

```encounter-table
name: Bael
creatures:
 - 1: Bael
```