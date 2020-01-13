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
