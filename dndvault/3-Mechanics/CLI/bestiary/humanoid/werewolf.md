---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/forest
- monster/environment/hill
- monster/size/medium
- monster/type/humanoid/human
- monster/type/humanoid/shapechanger
aliases: ["Werewolf"]
NoteIcon: monster
BestiaryType: humanoid (human, shapechanger)
SourceType: Bestiary
BookSource: Monster Manual p. 211. Available in the SRD and the Basic Rules.
---
# [Werewolf](3-Mechanics\CLI\bestiary\humanoid/werewolf.md)
*Source: Monster Manual p. 211. Available in the SRD and the Basic Rules.*  

A werewolf is a savage predator. In its humanoid form, a werewolf has heightened senses, a fiery temper, and a tendency to eat rare meat. Its wolf form is a fearsome predator, but its hybrid form is more terrifying by far-a furred and well-muscled humanoid body topped by a ravening wolf's head. A werewolf can wield weapons in hybrid form, though it prefers to tear foes apart with its powerful claws and bite.

Most werewolves flee civilized lands not long after becoming afflicted. Those that reject the curse fear what will happen if they remain among their friends and family. Those that embrace the curse fear discovery and the consequences of their murderous acts. In the wild, werewolves form packs that also include wolves and dire wolves.

> [!quote]- A quote from Thornstaff, elf druid  
> 
> The Company of the Black Moon—they used to be adventurers loyal to the realm. Now they roam the woods as a pack of werewolves. The king has promised estates, titles, and gold to anyone who can undo the curse afflicting them. I, for one, have no interest in such rewards.

## Lycanthropes

One of the most ancient and feared of all curses, lycanthropy can transform the most civilized humanoid into a ravening beast. In its natural humanoid form, a creature cursed by lycanthropy appears as its normal self. Over time, however, many lycanthropes acquire features suggestive of their animal form. In that animal form, a lycanthrope resembles a powerful version of a normal animal. On close inspection, its eyes show a faint spark of unnatural intelligence and might glow red in the dark.

Evil lycanthropes hide among normal folk, emerging in animal form at night to spread terror and bloodshed, especially under a full moon. Good lycanthropes are reclusive and uncomfortable around other civilized creatures, often living alone in wilderness areas far from villages and towns.

### Curse of Lycanthropy

A humanoid creature can be afflicted with the curse of lycanthropy after being wounded by a lycanthrope, or if one or both of its parents are lycanthropes. A [remove curse](/3-Mechanics/CLI/spells/remove-curse.md) spell can rid an afflicted lycanthrope of the curse, but a natural born lycanthrope can be freed of the curse only with a wish.

A lycanthrope can either resist its curse or embrace it. By resisting the curse, a lycanthrope retains its normal alignment and personality while in humanoid form. It lives its life as it always has, burying deep the bestial urges raging inside it. However, when the full moon rises, the curse becomes too strong to resist, transforming the individual into its beast form-or into a horrible hybrid form that combines animal and humanoid traits. When the moon wanes, the beast within can be controlled once again. Especially if the cursed creature is unaware of its condition, it might not remember the events of its transformation, though those memories often haunt a lycanthrope as bloody dreams.

Some individuals see little point in fighting the curse and accept what they are. With time and experience, they learn to master their shapechanging ability and can assume beast form or hybrid form at will. Most lycanthropes that embrace their bestial natures succumb to bloodlust, becoming evil, opportunistic creatures that prey on the weak.

> [!note] Player Characters as Lycanthropes
> 
> A character who becomes a lycanthrope retains his or her statistics except as specified by lycanthrope type. The character gains the lycanthrope's speeds in non-humanoid form, damage immunities, traits, and actions that don't involve equipment. The character is proficient with the lycanthrope's natural attacks, such as its bite or claws, which deal damage as shown in the lycanthrope's statistics. The character can't speak while in animal form.
> 
> A non-lycanthrope humanoid hit by an attack that carries the curse of lycanthropy must succeed on a Constitution saving throw (DC 8 + the lycanthrope's proficiency bonus + the lycanthrope's Constitution modifier) or be cursed. If the character embraces the curse, his or her alignment becomes the one defined for the lycanthrope. The DM is free to decide that a change in alignment places the character under DM control until the curse of lycanthropy is removed.
> 
> The following information applies to specific lycanthropes.
> 
> **.** **[Werebear](/3-Mechanics/CLI/bestiary/humanoid/werebear.md).** The character gains a Strength of 19 if his or her score isn't already higher, and a +1 bonus to AC while in bear or hybrid form (from natural armor). Attack and damage rolls for the natural weapons are based on Strength.
> 
> **[Wereboar](/3-Mechanics/CLI/bestiary/humanoid/wereboar.md).** The character gains a Strength of 17 if his or her score isn't already higher, and a +1 bonus to AC while in boar or hybrid form (from natural armor). Attack and damage rolls for the tusks are based on Strength. For the Charge trait, the DC is 8 + the character's proficiency bonus + Strength modifier.
> 
> **[Wererat](/3-Mechanics/CLI/bestiary/humanoid/wererat.md).** The character gains a Dexterity of 15 if his or her score isn't already higher. Attack and damage rolls for the bite are based on whichever is higher of the character's Strength and Dexterity.
> 
> **[Weretiger](/3-Mechanics/CLI/bestiary/humanoid/weretiger.md).** The character gains a Strength of 17 if his or her score isn't already higher. Attack and damage rolls for the natural weapons are based on Strength. For the Pounce trait, the DC is 8 + the character's proficiency bonus + Strength modifier.
> 
> **[Werewolf](/3-Mechanics/CLI/bestiary/humanoid/werewolf.md).** The character gains a Strength of 15 if his or her score isn't already higher, and a +1 bonus to AC while in wolf or hybrid form (from natural armor). Attack and damage rolls for the natural weapons are based on Strength.
^player-characters-as-lycanthropes

```statblock
"name": "Werewolf"
"size": "Medium"
"type": "humanoid"
"subtype": "human, shapechanger"
"alignment": "Chaotic Evil"
"ac": !!int "11"
"ac_class": "12 from natural armor in wolf or hybrid form"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "15"
- !!int "13"
- !!int "14"
- !!int "10"
- !!int "11"
- !!int "10"
"speed": "30 ft. (40 ft. in wolf form)"
"skillsaves":
  "Stealth": !!int "3"
  "Perception": !!int "4"
"damage_immunities": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't silvered"
"senses": "passive Perception 14"
"languages": "Common (can't speak in wolf form)"
"cr": "3"
"traits":
- "desc": "The werewolf can use its action to polymorph into a wolf-humanoid hybrid\
    \ or into a wolf, or back into its true form, which is humanoid. Its statistics,\
    \ other than its AC, are the same in each form. Any equipment it is wearing or\
    \ carrying isn't transformed. It reverts to its true form if it dies."
  "name": "Shapechanger"
- "desc": "The werewolf has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
"actions":
- "desc": "The werewolf makes two attacks: two with its spear (humanoid form) or one\
    \ with its bite and one with its claws (hybrid form)."
  "name": "Multiattack (Humanoid or Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage. If the target is\
    \ a humanoid, it must succeed on a DC 12 Constitution saving throw or be cursed\
    \ with werewolf lycanthropy."
  "name": "Bite (Wolf or Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d4 + 2|text(7) (2d4 + 2) slashing damage."
  "name": "Claws (Hybrid Form Only)"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 20/60 ft., one creature. Hit: dice:1d6 + 2|text(5) (1d6 + 2)\
    \ piercing damage, or dice:1d8 + 2|text(6) (1d8 + 2) piercing damage if used\
    \ with two hands to make a melee attack."
  "name": "Spear (Humanoid Form Only)"
"source":
- "MM"
- "CoS"
- "PotA"
- "SKT"
- "WDH"
- "GoS"
- "BGDIA"
- "ERLW"
- "IMR"
- "EGW"
- "IDRotF"
- "CM"
- "BMT"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Werewolf.webp"
```
^statblock

```encounter-table
name: Werewolf
creatures:
 - 1: Werewolf
```

## Environment

forest, hill