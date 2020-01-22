# Many real WPF Examples by MS (actually 203)

- https://github.com/Microsoft/WPF-Samples

# Paid Tutorials

- [With MSDN Subscription yout get some free Tutorials on Pluralsight and LinkedIn](https://my.visualstudio.com/benefits)
- [pluralsight.com/search?q=wpf](https://www.pluralsight.com/search?q=wpf&categories=course&sort=relevance)
- [linkedin.com/learning/windows-presentation-foundation-1](https://www.linkedin.com/learning/windows-presentation-foundation-1-build-dramatic-desktop-applications)
- https://www.udemy.com/topic/wpf/

# Free Tutorials

- [C # WPF UI Tutorials by AngelSix @YouTube](https://www.youtube.com/playlist?list=PLrW43fNmjaQVYF4zgsD0oL9Iv6u23PI6M)
- https://channel9.msdn.com/Tags/wpf

# WPF Debug
```
System.Diagnostics.PresentationTraceSources.DataBindingSource.Switch.Level = System.Diagnostics.SourceLevels.Critical;

, PresentationTraceSources.TraceLevel=High
```

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
