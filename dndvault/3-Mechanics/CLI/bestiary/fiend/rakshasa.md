---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/13
- monster/environment/urban
- monster/size/medium
- monster/type/fiend
aliases: ["Rakshasa"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Monster Manual p. 257, Tasha's Cauldron of Everything. Available in the SRD.
---
# [Rakshasa](3-Mechanics\CLI\bestiary\fiend/rakshasa.md)
*Source: Monster Manual p. 257, Tasha's Cauldron of Everything. Available in the SRD.*  

The rakshasa employs delicacy and misdirection in its pursuit of dominion over others. Few creatures ever see the fiend in its true form, for it can take on any guise it wants, although it prefers to masquerade as someone powerful or influential: a noble, cardinal, or rich merchant, for example. A rakshasa's true form combines the features of a human and a tiger, with one noteworthy deformity: its palms are where the backs of the hands would be on a human.

## Evil Spirits in Mortal Flesh

Rakshasas originated long ago in the Nine Hells, when powerful devils created a dark ritual to free their essence from their fiendish bodies in order to escape the Lower Planes. A rakshasa enters the Material Plane to feed its appetite for humanoid flesh and evil schemes. It selects its prey with care, taking pains to keep its presence in the world a secret.

## Evil Reborn

For a rakshasa, death on the Material Plane means an agonizing and torturous return to the Nine Hells, where its essence remains trapped until its body reforms-a process that can take months or years.

When the rakshasa is reborn, it has all the memories and knowledge of its former life, and it seeks retribution against the one who slew it. If the target has somehow slipped through its grasp, the rakshasa might punish its killer's family, friends, or descendants.

Like devils, rakshasas killed in the Nine Hells are forever destroyed.

> [!quote]- A quote from Rakshasa maxim  
> 
> Slay me once, shame on you.
> 
> Slay me twice, shame on me.


```statblock
"name": "Rakshasa"
"size": "Medium"
"type": "fiend"
"alignment": "Lawful Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "110"
"hit_dice": "13d8 + 52"
"stats":
- !!int "14"
- !!int "17"
- !!int "18"
- !!int "13"
- !!int "16"
- !!int "20"
"speed": "40 ft."
"skillsaves":
  "Deception": !!int "10"
  "Insight": !!int "8"
"damage_vulnerabilities": "piercing from magic weapons wielded by good creatures"
"damage_immunities": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Infernal"
"cr": "13"
"traits":
- "desc": "The rakshasa's innate spellcasting ability is Charisma (spell save DC 18,\
    \ dice: d20+10 (+10) to hit with spell attacks). The rakshasa can innately\
    \ cast the following spells, requiring no material components:\n\nAt will:\
    \ [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [disguise self](/3-Mechanics/CLI/spells/disguise-self.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\
    \n1/day each: [dominate person](/3-Mechanics/CLI/spells/dominate-person.md),\
    \ [fly](/3-Mechanics/CLI/spells/fly.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)\n\n3/day each: [charm\
    \ person](/3-Mechanics/CLI/spells/charm-person.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [major image](/3-Mechanics/CLI/spells/major-image.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Innate Spellcasting"
- "desc": "The rakshasa can't be affected or detected by spells of 6th level or lower\
    \ unless it wishes to be. It has advantage on saving throws against all other\
    \ spells and magical effects."
  "name": "Limited Magic Immunity"
"actions":
- "desc": "The rakshasa makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) slashing damage, and the target is\
    \ cursed if it is a creature. The magical curse takes effect whenever the target\
    \ takes a short or long rest, filling the target's thoughts with horrible images\
    \ and dreams. The cursed target gains no benefit from finishing a short or long\
    \ rest. The curse lasts until it is lifted by a [remove curse](/3-Mechanics/CLI/spells/remove-curse.md)\
    \ spell or similar magic."
  "name": "Claw"
"source":
- "MM"
- "TftYP"
- "WDMM"
- "GoS"
- "ERLW"
- "TCE"
- "CM"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Rakshasa.webp"
```
^statblock

```encounter-table
name: Rakshasa
creatures:
 - 1: Rakshasa
```

## Environment

urban