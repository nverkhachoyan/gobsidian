---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/10
- monster/environment/desert
- monster/environment/underdark
- monster/environment/urban
- monster/size/large
- monster/type/fiend/devil
aliases: ["Orthon"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 169
---
# [Orthon](3-Mechanics\CLI\bestiary\fiend/orthon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 169*  

## Orthon

When an archduke of the Nine Hells needs a creature tracked, found, and either done away with or captured, the task usually falls to an orthon. Orthons are infernal bounty hunters, tireless in their pursuit of their quarry across the multiverse.

## Unseen and All-Seeing

Orthons are infamous for their sharp senses. Because an orthon can become [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) at will, its quarry is often unaware of being hunted until the orthon strikes. The orthon's invisibility can be disrupted when the devil is attacked, however, so a strong counterattack is often the best defense against it.

## A Sporting Chance

Orthons value the challenge of the chase and the thrill of one-on-one combat above all else. An orthon's first loyalty is to its archduke, but if it has no immediate assignment, an orthon might work for anyone who offers it the promise of a worthy struggle against a lethal foe. Because they travel widely, orthons are unequaled as guides through the layers of the Nine Hells.

```statblock
"name": "Orthon (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "17"
"ac_class": "[half plate armor](/3-Mechanics/CLI/items/half-plate-armor.md)"
"hp": !!int "105"
"hit_dice": "10d10 + 50"
"stats":
- !!int "22"
- !!int "16"
- !!int "21"
- !!int "15"
- !!int "15"
- !!int "16"
"speed": "30 ft., climb 30 ft."
"saves":
  "Dexterity": !!int "7"
  "Wisdom": !!int "6"
  "Constitution": !!int "9"
"skillsaves":
  "Stealth": !!int "11"
  "Perception": !!int "10"
  "Survival": !!int "10"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., truesight 30 ft., passive Perception 20"
"languages": "Common, Infernal, telepathy 120 ft."
"cr": "10"
"traits":
- "desc": "The orthon can use a bonus action to become [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible).\
    \ Any equipment the orthon wears or carries is also [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ as long as the equipment is on its person. This invisibility ends immediately\
    \ after the orthon makes an attack roll or is hit by an attack."
  "name": "Invisibility Field"
- "desc": "The orthon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d4 + 6|text(11) (2d4 + 6) slashing damage, and the\
    \ target must make a DC 17 Constitution saving throw, taking dice:4d10|text(22)\
    \ (4d10) poison damage on a failed save, or half as much damage on a successful\
    \ one. On a failure, the target is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 minute. The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Infernal Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 100/400 ft.,\
    \ one target. Hit: dice:2d10 + 3|text(14) (2d10 + 3) piercing damage, plus\
    \ one of the following effects:\n\n- 1. Acid. The target must make a DC 17\
    \ Constitution saving throw, taking an additional dice:5d6|text(17) (5d6)\
    \ acid damage on a failed save, or half as much damage on a successful one.  \n\
    - 2. Blindness (1/Day). The target takes dice:1d10|text(5) (1d10) radiant\
    \ damage. In addition, the target and all other creatures within 20 feet of it\
    \ must each make a successful DC 17 Dexterity saving throw or be [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ until the end of the orthon's next turn.  \n- 3. Concussion. The target\
    \ and each creature within 20 feet of it must make a DC 17 Constitution saving\
    \ throw, taking dice:2d12|text(13) (2d12) thunder damage on a failed save,\
    \ or half as much damage on a successful one.  \n- 4. Entanglement. The target\
    \ must make a successful DC 17 Dexterity saving throw or be [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ for 1 hour by strands of sticky webbing. A [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ creature can escape by using an action to make a successful DC 17 Dexterity\
    \ or Strength check. Any creature other than an orthon that touches the [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ creature must make a successful DC 17 Dexterity saving throw or become similarly\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).  \n- 5. Paralysis\
    \ (1/Day). The target takes dice:4d10|text(22) (4d10) lightning damage and\
    \ must make a successful DC 17 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. The [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success.  \n- 6. Tracking. For the next 24 hours,\
    \ the orthon knows the direction and distance to the target, as long as it's on\
    \ the same plane of existence. If the target is on a different plane, the orthon\
    \ knows which one, but not the exact location there.  "
  "name": "Brass Crossbow"
"reactions":
- "desc": "When it is reduced to 15 hit points or fewer, the orthon causes itself\
    \ to explode. All other creatures within 30 feet of it must each make a DC 17\
    \ Dexterity saving throw, taking dice:2d8|text(9) (2d8) fire damage plus dice:2d8|text(9)\
    \ (2d8) thunder damage on a failed save, or half as much damage on a successful\
    \ one. This explosion destroys the orthon, its infernal dagger, and its brass\
    \ crossbow."
  "name": "Explosive Retribution"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Orthon.webp"
```
^statblock

```encounter-table
name: Orthon
creatures:
 - 1: Orthon
```

## Environment

desert, underdark, urban