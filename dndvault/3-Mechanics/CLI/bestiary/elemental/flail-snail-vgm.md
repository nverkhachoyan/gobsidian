---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/elemental
aliases: ["Flail Snail"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 144, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Flail Snail](3-Mechanics\CLI\bestiary\elemental/flail-snail-vgm.md)
*Source: Volo's Guide to Monsters p. 144, Dragon of Icespire Peak, Storm Lord's Wrath*  

A flail snail is a creature of elemental earth that is prized for its multi-hued shell. Hunters might be lulled into a false sense of confidence upon sighting this ponderous, seemingly nonhostile creature. If any other creature large enough to be a threat approaches too close, though, the snail unleashes a flash of scintillating light and then attacks with its mace-like tentacles.

## Trail of Treasure

Left undisturbed, a flail snail moves slowly along the ground, consuming everything on the surface, including rocks, sand, and soil, stopping to relish crystal growths and other large mineral deposits. It leaves behind a shimmering trail that quickly solidifies into a thin layer of a nearly transparent substance inedible to the snail. This glassy residue can be harvested and cut to form window panes of varying clearness. It can also be heated and spun into glass objects of other sorts. Some humanoids make a living from trailing flail snails to collect this glass.

> [!note] Using the Shell of a Flail Snail
> 
> A flail snail shell, which weighs about 250 pounds, has numerous uses. One intact shell can sell for 5,000 gp.
> 
> Many hunters seek the shell for its antimagic properties. A skilled armorer can make three shields from one shell. For 1 month, each shield gives its wielder the snail's Antimagic Shell trait. When the shield's magic fades, it leaves behind an exotic shield that is the perfect item from which to make a [spellguard shield](/3-Mechanics/CLI/items/spellguard-shield.md).
> 
> A flail snail shell can also be used to make a [robe of scintillating colors](/3-Mechanics/CLI/items/robe-of-scintillating-colors.md). The shell is ground and added to the dye while the garment is being fashioned. The powder is also a material component of the ritual that enchants the robe.
^using-the-shell-of-a-flail-snail

```statblock
"name": "Flail Snail (VGM)"
"size": "Large"
"type": "elemental"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "52"
"hit_dice": "5d10 + 25"
"stats":
- !!int "17"
- !!int "5"
- !!int "20"
- !!int "3"
- !!int "10"
- !!int "5"
"speed": "10 ft."
"damage_immunities": "fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 10"
"languages": ""
"cr": "3"
"traits":
- "desc": "The snail has advantage on saving throws against spells, and any creature\
    \ making a spell attack against the snail has disadvantage on the attack roll.\
    \ If the snail succeeds on its saving throw against a spell or a spell attack\
    \ misses it, an additional effect might occur, as determined by rolling a dice:\
    \ d6|avg|noform (d6):\n\n- 1–2. If the spell affects an area or has multiple\
    \ targets, it fails and has no effect. If the spell targets only the snail, it\
    \ has no effect on the snail and is reflected back at the caster, using the spell\
    \ slot level, spell save DC, attack bonus, and spellcasting ability of the caster.\
    \  \n- 3–4. No additional effect.  \n- 5–6. The snail's shell converts\
    \ some of the spell's energy into a burst of destructive force. Each creature\
    \ within 30 feet of the snail must make a DC 15 Constitution saving throw, taking\
    \ dice: 1d6|avg|noform (1d6) force damage per level of the spell on a failed\
    \ save, or half as much damage on a successful one.  "
  "name": "Antimagic Shell"
- "desc": "The flail snail has five flail tentacles. Whenever the snail takes 10 damage\
    \ or more on a single turn, one of its tentacles dies. If even one tentacle remains,\
    \ the snail regrows all dead ones within dice: 1d4|avg|noform (1d4) days.\
    \ If all its tentacles die, the snail retracts into its shell, gaining total cover,\
    \ and it begins wailing, a sound that can be heard for 600 feet, stopping only\
    \ when it dies dice: 5d6|avg|noform (5d6) minutes later. Healing magic that\
    \ restores limbs, such as the [regenerate](/3-Mechanics/CLI/spells/regenerate.md)\
    \ spell, can halt this dying process."
  "name": "Flail Tentacles"
"actions":
- "desc": "The flail snail makes as many flail tentacle attacks as it has flail tentacles,\
    \ all against the same target."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) bludgeoning damage."
  "name": "Flail Tentacle"
- "desc": "The snail's shell emits dazzling, colored light until the end of the snail's\
    \ next turn. During this time, the shell sheds bright light in a 30-foot radius\
    \ and dim light for an additional 30 feet, and creatures that can see the snail\
    \ have disadvantage on attack rolls against it. In addition, any creature within\
    \ the bright light and able to see the snail when this power is activated must\
    \ succeed on a DC 15 Wisdom saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the light ends."
  "name": "Scintillating Shell (Recharges after a Short or Long Rest)"
- "desc": "The flail snail withdraws into its shell, gaining a +4 bonus to AC until\
    \ it emerges. It can emerge from its shell as a bonus action on its turn."
  "name": "Shell Defense"
"source":
- "VGM"
- "ToA"
- "DIP"
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Flail%20Snail.webp"
```
^statblock

```encounter-table
name: Flail Snail
creatures:
 - 1: Flail Snail
```

## Environment

underdark, forest, swamp