---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/12
- monster/environment/arctic
- monster/environment/coastal
- monster/size/huge
- monster/type/giant/frost-giant
aliases: ["Frost Giant Everlasting One"]
NoteIcon: monster
BestiaryType: giant (frost giant)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 148
---
# [Frost Giant Everlasting One](3-Mechanics\CLI\bestiary\giant/frost-giant-everlasting-one-vgm.md)
*Source: Volo's Guide to Monsters p. 148*  

To hold its place or rise within the ordning, a frost giant must routinely face mighty foes in single combat. Some seek out magic that will aid them, but enchanted objects can be taken or lost. True greatness relies on personal prowess. Faced with this truth, a frost giant might seek a supernatural gift from Vaprak the Destroyer.

## Troll Eater

Frost giants mainly turn to Vaprak, a rapacious god of strength and hunger worshiped by ogres and trolls, out of desperation. Vaprak likes to tempt frost giants with dreams of glory followed by nightmares of bloody cannibalism. Those who don't shrink from such visions or report them to priests of Thrym receive more of the same. If a frost giant comes to relish these dreams and nightmares, as some do, Vaprak sets a troll upon a sacred quest to find the frost giant and meet it in secret. The troll offers up its own body to be devoured in Vaprak's name. Only the boldest and most determined frost giants can finish such a gory feast.

## Vaprak's Blessing

After devouring the troll sent by Vaprak, bones and all, a frost giant becomes an everlasting one, gaining tremendous strength, an ill temper, and a troll's regenerative ability. With these gifts, the frost giant can swiftly claim the title of jarl and easily fend off rivals for decades. However, if the frost giant doesn't give enough honor to Vaprak or fails to heed Vaprak's visions, injuries the frost giant sustains heal wrong, often resulting in discolored skin, warty scars, and vestigial body parts, such as extra digits, limbs, and even extra heads. The touch of Vaprak can no longer be hidden then, and the everlasting one is either killed or exiled by its clan. Sometimes small communities of everlasting ones gather and even reproduce, passing the "blessing" and worship of Vaprak from one generation to the next.

```statblock
"name": "Frost Giant Everlasting One (VGM)"
"size": "Huge"
"type": "giant"
"subtype": "frost giant"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "patchwork armor"
"hp": !!int "189"
"hit_dice": "14d12 + 98"
"stats":
- !!int "25"
- !!int "9"
- !!int "24"
- !!int "9"
- !!int "10"
- !!int "12"
"speed": "40 ft."
"saves":
  "Wisdom": !!int "4"
  "Strength": !!int "11"
  "Constitution": !!int "11"
"skillsaves":
  "Athletics": !!int "11"
  "Perception": !!int "4"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Giant"
"cr": "12"
"traits":
- "desc": "The giant has a 25% chance chance of having more than one head. If it has\
    \ more than one, it has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks and on saving throws against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Extra Heads"
- "desc": "The giant regains 10 hit points at the start of its turn. If the giant\
    \ takes acid or fire damage, this trait doesn't function at the start of its next\
    \ turn. The giant dies only if it starts its turn with 0 hit points and doesn't\
    \ regenerate."
  "name": "Regeneration"
- "desc": "As a bonus action, the giant can enter a rage at the start of its turn.\
    \ The rage lasts for 1 minute or until the giant is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\
    \ While raging, the giant gains the following benefits:\n\n- The giant has advantage\
    \ on Strength checks and Strength saving throws\n\n- When it makes a melee weapon\
    \ attack, the giant gains a +4 bonus to the damage roll.\n\n- The giant has resistance\
    \ to bludgeoning, piercing, and slashing damage."
  "name": "Vaprak's Rage (Recharges after a Short or Long Rest)"
"actions":
- "desc": "The giant makes two attacks with its greataxe."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d12 + 7|text(26) (3d12 + 7) slashing damage, or dice:3d12\
    \ + 11|text(30) (3d12 + 11) slashing damage while raging."
  "name": "Greataxe"
- "desc": "Ranged Weapon Attack: dice: d20+11 (+11) to hit, range 60/240 ft.,\
    \ one target. Hit: dice:4d10 + 7|text(29) (4d10 + 7) bludgeoning damage."
  "name": "Rock"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Frost%20Giant%20Everlasting%20One.webp"
```
^statblock

```encounter-table
name: Frost Giant Everlasting One
creatures:
 - 1: Frost Giant Everlasting One
```

## Environment

coastal, arctic