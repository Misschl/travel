# Generated by Django 2.2 on 2020-06-19 12:58

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('user', '0002_auto_20200619_2056'),
    ]

    operations = [
        migrations.AlterField(
            model_name='user',
            name='sex',
            field=models.SmallIntegerField(blank=True, choices=[(1, '男'), (0, '女')], null=True),
        ),
    ]