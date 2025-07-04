---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/5
- monster/environment/arctic
- monster/environment/forest
- monster/environment/hill
- monster/size/medium
- monster/type/humanoid/human
- monster/type/humanoid/shapechanger
aliases: ["Werebear"]
NoteIcon: monster
BestiaryType: humanoid (human, shapechanger)
SourceType: Bestiary
BookSource: Monster Manual p. 208. Available in the SRD.
---
# [Werebear](3-Mechanics\CLI\bestiary\humanoid/werebear.md)
*Source: Monster Manual p. 208. Available in the SRD.*  

Werebears are powerful lycanthropes with the ability to temper their monstrous natures and reject their violent impulses. In humanoid form, they are large, muscular, and covered in hair matching the color of their ursine form's fur. A werebear is a loner by nature, fearing what might happen to innocent creatures around it when its bestial nature takes over.

When a werebear transforms, it grows to enormous size, lashing out with weapons or claws. It fights with the ferocity of a bear, though even in its bestial forms, it avoids biting so as to not pass on its curse. Typically, a werebear passes on its lycanthropy only to chosen companions or apprentices, spending the time that follows helping the new lycanthrope accept the curse in order to control it.

Solitary creatures, werebears act as wardens over their territory, protecting flora and fauna alike from humanoid or monstrous intrusion. Though most werebears are of good alignment, some are every bit as evil as other lycanthropes.

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
"name": "Werebear"
"size": "Medium"
"type": "humanoid"
"subtype": "human, shapechanger"
"alignment": "Neutral Good"
"ac": !!int "10"
"ac_class": "11 from natural armor in bear or hybrid form"
"hp": !!int "135"
"hit_dice": "18d8 + 54"
"stats":
- !!int "19"
- !!int "10"
- !!int "17"
- !!int "11"
- !!int "12"
- !!int "12"
"speed": "30 ft. (40 ft., climb 30 ft. in bear or hybrid form)"
"skillsaves":
  "Perception": !!int "7"
"damage_immunities": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't silvered"
"senses": "passive Perception 17"
"languages": "Common (can't speak in bear form)"
"cr": "5"
"traits":
- "desc": "The werebear can use its action to polymorph into a Large bear-humanoid\
    \ hybrid or into a Large bear, or back into its true form, which is humanoid.\
    \ Its statistics, other than its size and AC, are the same in each form. Any equipment\
    \ it is wearing or carrying isn't transformed. It reverts to its true form if\
    \ it dies."
  "name": "Shapechanger"
- "desc": "The werebear has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
"actions":
- "desc": "In bear form, the werebear makes two claw attacks. In humanoid form, it\
    \ makes two greataxe attacks. In hybrid form, it can attack like a bear or a humanoid."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 4|text(15) (2d10 + 4) piercing damage. If the target\
    \ is a humanoid, it must succeed on a DC 14 Constitution saving throw or be cursed\
    \ with werebear lycanthropy."
  "name": "Bite (Bear or Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) slashing damage."
  "name": "Claw (Bear or Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d12 + 4|text(10) (1d12 + 4) slashing damage."
  "name": "Greataxe (Humanoid or Hybrid Form Only)"
"source":
- "MM"
- "GoS"
- "EGW"
- "KftGV"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Werebear.webp"
```
^statblock

```encounter-table
name: Werebear
creatures:
 - 1: Werebear
```

## Environment

forest, hill, arctic