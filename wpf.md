# Many real WPF Examples by MS

- WPF for .Net Core 3.1: https://github.com/microsoft/WPF-Samples/tree/master
- WPF for .Net 4.7.2: https://github.com/microsoft/WPF-Samples/tree/netframework

# Paid Tutorials

- [With MSDN Subscription yout get some free Tutorials on Pluralsight and LinkedIn](https://my.visualstudio.com/benefits)
- [pluralsight.com/search?q=wpf](https://www.pluralsight.com/search?q=wpf&categories=course&sort=relevance)
- [linkedin.com/learning/windows-presentation-foundation-1](https://www.linkedin.com/learning/windows-presentation-foundation-1-build-dramatic-desktop-applications)
- https://www.udemy.com/topic/wpf/

# Free Tutorials

- [C # WPF UI Tutorials by AngelSix @YouTube](https://www.youtube.com/playlist?list=PLrW43fNmjaQVYF4zgsD0oL9Iv6u23PI6M)
- https://channel9.msdn.com/Tags/wpf

# Tools

- https://github.com/snoopwpf/snoopwpf

# Binding 
```
# Attached Property
{Binding Path=(helpers:WpfHelpers.CustomDataContext).IsAnyUniversalSafetyRelevant}"
```

# Binding Debugging
 * `PresentationTraceSources.TraceLevel=High` (`Text={Binding Path=Foobar, PresentationTraceSources.TraceLevel=High}`)
 * `PresentationTraceSources.SetTraceLevel(binding, PresentationTraceLevel.High);`
  * ```
var binding = new Binding("Foobar");
PresentationTraceSources.SetTraceLevel(binding, PresentationTraceLevel.High); // important: must happen before "SetBinding()"
\_target.SetBinding(TextBox.TextProperty, binding);
  ```
`System.Diagnostics.PresentationTraceSources.DataBindingSource.Switch.Level = System.Diagnostics.SourceLevels.Critical;`

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
