# Many real WPF Examples by MS

- WPF for .Net Core 3.1: https://github.com/microsoft/WPF-Samples/tree/master
- WPF for .Net 4.7.2: https://github.com/microsoft/WPF-Samples/tree/netframework

# Paid Tutorials

- [With MSDN Subscription yout get some free Tutorials on Pluralsight and LinkedIn](https://my.visualstudio.com/benefits)
- [pluralsight.com/search?q=wpf](https://www.pluralsight.com/search?q=wpf&categories=course&sort=relevance)
- [linkedin.com/learning/windows-presentation-foundation-1](https://www.linkedin.com/learning/windows-presentation-foundation-1-build-dramatic-desktop-applications)
- https://www.udemy.com/topic/wpf/

# Free Tutorials

- [Alles, was Sie über XAML wissen müssen | Thomas Claudius Huber](https://youtu.be/LgqeifxPCuk)
- [Blend Tips: Intro to Using Behaviors - Part 1](https://youtu.be/1sn-uh2yPCo)
- [C # WPF UI Tutorials by AngelSix @YouTube](https://www.youtube.com/playlist?list=PLrW43fNmjaQVYF4zgsD0oL9Iv6u23PI6M)
- https://channel9.msdn.com/Tags/wpf
- [WPF Path Markup](https://www.youtube.com/watch?v=AjuUiu5y-sk)

# Tools

- https://github.com/snoopwpf/snoopwpf

# Libraries

- https://github.com/Infragistics/InfragisticsThemesForMicrosoftControls

# ContentPresenter vs ContentControl

- https://stackoverflow.com/questions/1287995/whats-the-difference-between-contentcontrol-and-contentpresenter

# Binding 

[MSDN: Binding Path Syntax](https://docs.microsoft.com/en-us/dotnet/framework/wpf/data/binding-declarations-overview?redirectedfrom=MSDN#Path_Syntax)

```
# Attached Property
{Binding Path=(myLib:MyClass.MyAttachedProperty).MyViewModelProperty}"
```

# inner Border-collaps Table
https://stackoverflow.com/questions/4572631/is-it-possible-to-emulate-border-collapse-ala-css-in-a-wpf-itemscontrol
```
<ItemsControl ItemsSource="{Binding MyList}">
<ItemsControl.ItemTemplate>
    <DataTemplate>
	<Grid x:Name="grid" Margin="0,8,0,0" />
	
	<DataTemplate.Triggers>
	    <DataTrigger Binding="{Binding RelativeSource={RelativeSource Mode=PreviousData}}" Value="{x:Null}">
		<Setter TargetName="grid" Property="Margin" Value="0,0,0,0" />
	    </DataTrigger>
	</DataTemplate.Triggers>
    </DataTemplate>
</ItemsControl.ItemTemplate>
</ItemsControl>
```

# Binding Debugging
 * `PresentationTraceSources.TraceLevel=High`
 	* `<TextBox Text="{Binding Path=Foobar, PresentationTraceSources.TraceLevel=High}" />`
 * `PresentationTraceSources.SetTraceLevel(binding, PresentationTraceLevel.High);`
     * full example:
    ```
        var binding = new Binding("Foobar");
        PresentationTraceSources.SetTraceLevel(binding, PresentationTraceLevel.High); // MUST happen before "SetBinding()"
        \_target.SetBinding(TextBox.TextProperty, binding);
    ```

 * `System.Diagnostics.PresentationTraceSources.DataBindingSource.Switch.Level = System.Diagnostics.SourceLevels.Critical;`

# WPF DesignTime

```
xmlns:ComponentModel="clr-namespace:System.ComponentModel;assembly=PresentationFramework"

<DataTrigger Binding="{Binding (ComponentModel:DesignerProperties.IsInDesignMode), RelativeSource={RelativeSource Self}}" Value="True">
	<Setter Property="Visibility" Value="Visible" /> </DataTrigger>

<ResourceDictionary>
	<ResourceDictionary.MergedDictionaries>
		<ResourceDictionary Source="pack://application:,,,/MyApplication.View;component/Resources.xaml"/>
	</ResourceDictionary.MergedDictionaries>
</ResourceDictionary>
```

# WPF Style BasedOn Default Style
```
<Style TargetType="{x:Type igWPF:XamComboEditor}" BasedOn="{StaticResource {x:Type igWPF:XamComboEditor}}">
  <Setter Property="Format" Value="N0" /> 
</Style>
```

# WPF Software-Renderer
```
System.Windows.Media.RenderOptions.ProcessRenderMode = System.Windows.Interop.RenderMode.SoftwareOnly;
```

# Binding


## Binding Directions
![wpf binding directions](https://github.com/juwens/cheat-sheet/raw/master/assets/wpf_binding.png)

# List UI Synchronization Lock	
```
BindingOperations.EnableCollectionSynchronization(MyList, _myListLock);
```
